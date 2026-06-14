package tui

import "github.com/charmbracelet/lipgloss"

// Theme holds the colors used across the UI.
type Theme struct {
	Name      string
	Primary   lipgloss.Color
	Secondary lipgloss.Color
	Accent    lipgloss.Color
	Muted     lipgloss.Color
	Success   lipgloss.Color
	Warning   lipgloss.Color
	Error     lipgloss.Color
	Text      lipgloss.Color
}

var themes = map[string]Theme{
	"catppuccin": {
		Name: "Catppuccin Mocha", Primary: "#a6e3a1", Secondary: "#cba6f7",
		Accent: "#f38ba8", Muted: "#6c7086", Success: "#a6e3a1",
		Warning: "#f9e2af", Error: "#f38ba8", Text: "#cdd6f4",
	},
	"dracula": {
		Name: "Dracula", Primary: "#50fa7b", Secondary: "#bd93f9",
		Accent: "#ff79c6", Muted: "#6272a4", Success: "#50fa7b",
		Warning: "#f1fa8c", Error: "#ff5555", Text: "#f8f8f2",
	},
	"nord": {
		Name: "Nord", Primary: "#88c0d0", Secondary: "#81a1c1",
		Accent: "#b48ead", Muted: "#4c566a", Success: "#a3be8c",
		Warning: "#ebcb8b", Error: "#bf616a", Text: "#eceff4",
	},
	"tokyonight": {
		Name: "Tokyo Night", Primary: "#7aa2f7", Secondary: "#bb9af7",
		Accent: "#f7768e", Muted: "#565f89", Success: "#9ece6a",
		Warning: "#e0af68", Error: "#f7768e", Text: "#c0caf5",
	},
	"neon": {
		Name: "Neon", Primary: "#00ff9d", Secondary: "#7c3aed",
		Accent: "#ff6b6b", Muted: "#64748b", Success: "#00ff9d",
		Warning: "#fbbf24", Error: "#ff6b6b", Text: "#e2e8f0",
	},
}

func themeByName(name string) Theme {
	if t, ok := themes[name]; ok {
		return t
	}
	return themes["catppuccin"]
}

func themeNames() []string {
	return []string{"catppuccin", "dracula", "nord", "tokyonight", "neon"}
}

// styles bundles pre-built lipgloss styles for a theme.
type styles struct {
	header   lipgloss.Style
	logo     lipgloss.Style
	user     lipgloss.Style
	agent    lipgloss.Style
	tool     lipgloss.Style
	ok       lipgloss.Style
	warn     lipgloss.Style
	err      lipgloss.Style
	muted    lipgloss.Style
	status   lipgloss.Style
	box      lipgloss.Style
	menu     lipgloss.Style
	selected lipgloss.Style
	prompt   lipgloss.Style
}

func newStyles(t Theme) styles {
	return styles{
		header:   lipgloss.NewStyle().Foreground(t.Primary).Bold(true),
		logo:     lipgloss.NewStyle().Foreground(t.Primary).Bold(true),
		user:     lipgloss.NewStyle().Foreground(t.Secondary).Bold(true),
		agent:    lipgloss.NewStyle().Foreground(t.Text),
		tool:     lipgloss.NewStyle().Foreground(t.Accent),
		ok:       lipgloss.NewStyle().Foreground(t.Success),
		warn:     lipgloss.NewStyle().Foreground(t.Warning),
		err:      lipgloss.NewStyle().Foreground(t.Error).Bold(true),
		muted:    lipgloss.NewStyle().Foreground(t.Muted),
		status:   lipgloss.NewStyle().Foreground(t.Muted),
		box:      lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(t.Muted).Padding(0, 1),
		menu:     lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(t.Secondary).Padding(0, 1),
		selected: lipgloss.NewStyle().Foreground(t.Primary).Bold(true),
		prompt:   lipgloss.NewStyle().Foreground(t.Primary).Bold(true),
	}
}
