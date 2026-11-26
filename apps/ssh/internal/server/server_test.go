package server

import (
	"bytes"
	"io"
	"net"
	"path/filepath"
	"strings"
	"testing"
	"time"

	cryptossh "golang.org/x/crypto/ssh"
)

func TestSSHSessionRendersPortfolio(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	defer ln.Close()

	cfg := Config{
		Addr:        ln.Addr().String(),
		HostKeyPath: filepath.Join(t.TempDir(), "host_ed25519"),
	}
	s, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	errCh := make(chan error, 1)
	go func() {
		errCh <- ServeListener(s, ln)
	}()

	clientCfg := &cryptossh.ClientConfig{
		User:            "visitor",
		Auth:            nil,
		HostKeyCallback: cryptossh.InsecureIgnoreHostKey(),
		Timeout:         2 * time.Second,
	}
	var client *cryptossh.Client
	var dialErr error
	deadline := time.Now().Add(5 * time.Second)
	for time.Now().Before(deadline) {
		client, dialErr = cryptossh.Dial("tcp", ln.Addr().String(), clientCfg)
		if dialErr == nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	if dialErr != nil {
		t.Fatalf("dial: %v", dialErr)
	}
	defer client.Close()

	sess, err := client.NewSession()
	if err != nil {
		t.Fatal(err)
	}
	defer sess.Close()

	var out bytes.Buffer
	sess.Stdout = &out
	sess.Stderr = &out
	modes := cryptossh.TerminalModes{
		cryptossh.ECHO:          0,
		cryptossh.TTY_OP_ISPEED: 14400,
		cryptossh.TTY_OP_OSPEED: 14400,
	}
	if err := sess.RequestPty("xterm-256color", 40, 100, modes); err != nil {
		t.Fatalf("pty: %v", err)
	}
	stdin, err := sess.StdinPipe()
	if err != nil {
		t.Fatal(err)
	}
	if err := sess.Shell(); err != nil {
		t.Fatalf("shell: %v", err)
	}
	// Answer OSC 11 if the program still queries the terminal.
	_, _ = io.WriteString(stdin, "\x1b]11;rgb:0a0a/0a0a/0a0a\x07")

	readDeadline := time.Now().Add(5 * time.Second)
	var body string
	for time.Now().Before(readDeadline) {
		body = stripANSI(out.String())
		if strings.Contains(body, "Brendan Palkowski") {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
	if !strings.Contains(body, "Brendan Palkowski") {
		t.Fatalf("session output missing name:\n%s", body)
	}
	if !strings.Contains(body, "About") {
		t.Fatalf("session output missing nav:\n%s", body)
	}

	_, _ = io.WriteString(stdin, "q")
	_ = sess.Wait()
}

func stripANSI(s string) string {
	var b strings.Builder
	i := 0
	for i < len(s) {
		if s[i] == 0x1b {
			if i+1 < len(s) && s[i+1] == '[' {
				j := i + 2
				for j < len(s) && (s[j] < '@' || s[j] > '~') {
					j++
				}
				if j < len(s) {
					i = j + 1
					continue
				}
			}
			if i+1 < len(s) && s[i+1] == ']' {
				if j := strings.Index(s[i:], "\x1b\\"); j != -1 {
					i += j + 2
					continue
				}
			}
		}
		b.WriteByte(s[i])
		i++
	}
	return b.String()
}
