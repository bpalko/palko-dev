package tui

import "github.com/charmbracelet/lipgloss"

// Palette matches apps/web (black / green, Michigan gold on About).
const (
	colorText     = "#F5F5F5"
	colorDim      = "#B3B3B3"
	colorMuted    = "#808080"
	colorAccent   = "#4ADE80"
	colorBorder   = "#2A2A2A"
	colorGold     = "#FFB347"
	colorMaize    = "#FFCB05"
	colorMichigan = "#00274C"
)

type styles struct {
	title    lipgloss.Style
	role     lipgloss.Style
	dim      lipgloss.Style
	muted    lipgloss.Style
	accent   lipgloss.Style
	heading  lipgloss.Style
	selected lipgloss.Style
	tab      lipgloss.Style
	help     lipgloss.Style
	gold     lipgloss.Style
	maize    lipgloss.Style
	box      lipgloss.Style
	status   lipgloss.Style
}

func newStyles(r *lipgloss.Renderer) styles {
	n := func() lipgloss.Style {
		if r == nil {
			return lipgloss.NewStyle()
		}
		return r.NewStyle()
	}
	return styles{
		title:  n().Bold(true).Foreground(lipgloss.Color(colorText)),
		role:   n().Foreground(lipgloss.Color(colorAccent)),
		dim:    n().Foreground(lipgloss.Color(colorDim)),
		muted:  n().Foreground(lipgloss.Color(colorMuted)),
		accent: n().Foreground(lipgloss.Color(colorAccent)),
		heading: n().Bold(true).Foreground(lipgloss.Color(colorText)).
			BorderBottom(true).
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(colorBorder)),
		selected: n().Bold(true).Foreground(lipgloss.Color("#0A0A0A")).
			Background(lipgloss.Color(colorAccent)).
			Padding(0, 1),
		tab:  n().Foreground(lipgloss.Color(colorDim)).Padding(0, 1),
		help: n().Foreground(lipgloss.Color(colorMuted)),
		gold: n().Foreground(lipgloss.Color(colorGold)),
		maize: n().Bold(true).Foreground(lipgloss.Color(colorMaize)).
			Background(lipgloss.Color(colorMichigan)).
			Padding(0, 1),
		box: n().Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(colorAccent)).
			Padding(0, 1),
		status: n().Foreground(lipgloss.Color(colorDim)),
	}
}
