package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"

	"github.com/bpalko/palko-dev/ssh/internal/content"
)

type section int

const (
	sectionAbout section = iota
	sectionExperience
	sectionBooks
	sectionBar
	sectionMusic
	sectionBlog
	sectionContact
)

var sections = []string{
	"About",
	"Experience",
	"Bookshelf",
	"Bar Cart",
	"Music",
	"Blog",
	"Contact",
}

type Model struct {
	width   int
	height  int
	section section
	styles  styles
	vp      viewport.Model
}

func New(width, height int, r *lipgloss.Renderer) Model {
	if width <= 0 {
		width = 80
	}
	if height <= 0 {
		height = 24
	}
	m := Model{
		width:  width,
		height: height,
		styles: newStyles(r),
		vp:     viewport.New(maxInt(width-2, 20), maxInt(height-10, 5)),
	}
	m.syncViewport()
	return m
}

func FromSession(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	m := New(pty.Window.Width, pty.Window.Height, nil)
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.syncViewport()
		return m, nil

	case tea.KeyMsg:
		if msg.Type == tea.KeyRunes && len(msg.Runes) > 1 {
			var next tea.Model = m
			var cmds []tea.Cmd
			for _, r := range msg.Runes {
				var cmd tea.Cmd
				next, cmd = next.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{r}})
				cmds = append(cmds, cmd)
			}
			return next, tea.Batch(cmds...)
		}

		switch msg.String() {
		case "q", "ctrl+c", "ctrl+d":
			return m, tea.Quit
		case "tab", "l", "right":
			m.section = (m.section + 1) % section(len(sections))
			m.syncViewport()
			m.vp.GotoTop()
			return m, nil
		case "shift+tab", "h", "left":
			m.section = (m.section - 1 + section(len(sections))) % section(len(sections))
			m.syncViewport()
			m.vp.GotoTop()
			return m, nil
		case "1", "2", "3", "4", "5", "6", "7":
			m.section = section(msg.String()[0] - '1')
			m.syncViewport()
			m.vp.GotoTop()
			return m, nil
		case "g", "home":
			m.vp.GotoTop()
			return m, nil
		case "G", "end":
			m.vp.GotoBottom()
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.vp, cmd = m.vp.Update(msg)
	return m, cmd
}

func (m *Model) syncViewport() {
	header := lipgloss.Height(m.headerView())
	nav := lipgloss.Height(m.navView())
	footer := lipgloss.Height(m.footerView())
	h := m.height - header - nav - footer - 1
	if h < 3 {
		h = 3
	}
	w := m.width - 2
	if w < 20 {
		w = 20
	}
	m.vp.Width = w
	m.vp.Height = h
	m.vp.SetContent(m.sectionContent())
}

func (m Model) View() string {
	if m.width < 1 || m.height < 1 {
		return "…"
	}
	return lipgloss.JoinVertical(lipgloss.Left,
		m.headerView(),
		m.navView(),
		m.vp.View(),
		m.footerView(),
	)
}

func (m Model) headerView() string {
	st := m.styles
	w := maxInt(m.width-2, 20)
	title := st.title.Render(content.Name)
	role := st.role.Render("$ " + content.Role)
	meta := st.dim.Render(fmt.Sprintf("%s · %s · %s", content.Location, content.Age, "palko.dev"))
	web := st.muted.Render("web hosted on Cloudflare")

	if m.width < 56 {
		return lipgloss.JoinVertical(lipgloss.Left, title, role, meta)
	}

	left := lipgloss.JoinVertical(lipgloss.Left, title, role, meta)
	right := lipgloss.JoinVertical(lipgloss.Right, st.accent.Render("palko.dev"), web)
	row := lipgloss.JoinHorizontal(lipgloss.Top,
		left,
		strings.Repeat(" ", maxInt(2, w-lipgloss.Width(left)-lipgloss.Width(right)-2)),
		right,
	)
	return st.box.Width(w).Render(row)
}

func (m Model) navView() string {
	st := m.styles
	var parts []string
	for i, name := range sections {
		label := fmt.Sprintf("%d:%s", i+1, name)
		if section(i) == m.section {
			parts = append(parts, st.selected.Render(label))
		} else {
			parts = append(parts, st.tab.Render(label))
		}
	}
	row := strings.Join(parts, "")
	if lipgloss.Width(row) > m.width {
		return st.selected.Render(fmt.Sprintf("%d:%s", int(m.section)+1, sections[m.section]))
	}
	return row
}

func (m Model) footerView() string {
	help := "1-7 tabs  h/l or ←/→  j/k scroll  g/G top/end  q quit"
	if m.width < 64 {
		help = "1-7  h/l  j/k  q"
	}
	scroll := ""
	if m.vp.TotalLineCount() > m.vp.Height {
		scroll = fmt.Sprintf("  %3.0f%%", m.vp.ScrollPercent()*100)
	}
	return m.styles.help.Render(help + scroll)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
