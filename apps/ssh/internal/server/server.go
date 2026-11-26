package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	wishtea "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/muesli/termenv"

	"github.com/bpalko/palko-dev/ssh/internal/tui"
)

const (
	defaultAddr        = "0.0.0.0:2222"
	defaultHostKeyPath = ".ssh/host_ed25519"
	shutdownTimeout    = 30 * time.Second
	idleTimeout        = 10 * time.Minute
)

type Config struct {
	Addr        string
	HostKeyPath string
}

func FromEnv() Config {
	return Config{
		Addr:        envOr("SSH_ADDR", defaultAddr),
		HostKeyPath: envOr("HOST_KEY_PATH", defaultHostKeyPath),
	}
}

func New(cfg Config) (*ssh.Server, error) {
	if cfg.Addr == "" {
		cfg.Addr = defaultAddr
	}
	if cfg.HostKeyPath == "" {
		cfg.HostKeyPath = defaultHostKeyPath
	}
	if err := os.MkdirAll(filepath.Dir(cfg.HostKeyPath), 0o700); err != nil {
		return nil, fmt.Errorf("host key dir: %w", err)
	}

	// SSH sessions cannot answer OSC color queries; pin a profile up front.
	lipgloss.SetColorProfile(termenv.ANSI256)

	return wish.NewServer(
		wish.WithAddress(cfg.Addr),
		wish.WithHostKeyPath(cfg.HostKeyPath),
		wish.WithIdleTimeout(idleTimeout),
		wish.WithMiddleware(
			wishtea.Middleware(tui.FromSession),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
}

func Run(cfg Config) error {
	s, err := New(cfg)
	if err != nil {
		return err
	}
	return Serve(s)
}

func Serve(s *ssh.Server) error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Starting SSH portfolio", "addr", s.Addr)
	go func() {
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
			log.Error("SSH server failed", "error", err)
			done <- syscall.SIGTERM
		}
	}()

	<-done
	log.Info("Stopping SSH portfolio")
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		return err
	}
	return nil
}

func ServeListener(s *ssh.Server, ln net.Listener) error {
	return s.Serve(ln)
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
