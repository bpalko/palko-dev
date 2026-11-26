package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestAboutRendersBrendan(t *testing.T) {
	m := New(100, 40, nil)
	got := stripANSI(m.View())
	for _, want := range []string{"Brendan Palkowski", "About", "First American", "Cloudflare"} {
		if !strings.Contains(got, want) {
			t.Fatalf("about view missing %q\n%s", want, got)
		}
	}
}

func TestNumberKeysSwitchSections(t *testing.T) {
	m := New(100, 40, nil)
	cases := []struct {
		key  string
		want string
	}{
		{"2", "Senior Software Engineer"},
		{"3", "Kitchen Confidential"},
		{"4", "Buffalo Trace"},
		{"5", "Yo Daddy's Yacht Rock"},
		{"6", "Come Back Later"},
		{"7", "github.com/bpalko"},
		{"1", "University of Michigan"},
	}
	var model tea.Model = m
	for _, tc := range cases {
		next, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(tc.key)})
		model = next
		got := stripANSI(model.(Model).View())
		if !strings.Contains(got, tc.want) {
			t.Fatalf("key %s: missing %q\n%s", tc.key, tc.want, got)
		}
	}
}

func TestTabAndQuit(t *testing.T) {
	m := New(80, 24, nil)
	next, _ := m.Update(tea.KeyMsg{Type: tea.KeyTab})
	got := stripANSI(next.(Model).View())
	if !strings.Contains(got, "First American") && !strings.Contains(got, "Experience") {
		t.Fatalf("tab should move to experience:\n%s", got)
	}
	_, cmd := next.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'q'}})
	if cmd == nil {
		t.Fatal("q should quit")
	}
	_, cmd = next.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Fatal("ctrl+c should quit")
	}
}

func TestNarrowTerminalStillRenders(t *testing.T) {
	m := New(40, 16, nil)
	next, _ := m.Update(tea.WindowSizeMsg{Width: 40, Height: 12})
	got := stripANSI(next.(Model).View())
	if !strings.Contains(got, "Brendan Palkowski") {
		t.Fatalf("narrow view missing name:\n%s", got)
	}
	if strings.Contains(got, "panic") {
		t.Fatal(got)
	}
}

func TestWrapRespectsWidth(t *testing.T) {
	s := wrapText("one two three four five six seven eight", 12)
	for _, line := range strings.Split(s, "\n") {
		if len(line) > 12 {
			t.Fatalf("line too long (%d): %q", len(line), line)
		}
	}
}
