package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/rozaq/termux-vibe-coding/internal/config"
	"github.com/rozaq/termux-vibe-coding/internal/tui"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal membaca konfigurasi: %v\n", err)
		os.Exit(1)
	}

	// Quick-cmds: if user passes flags, handle them before starting TUI
	// (for now just launch TUI)

	m, err := tui.NewModel(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal memulai: %v\n", err)
		fmt.Fprintln(os.Stderr, "Tips: export OPENROUTER_API_KEY=sk-or-... dan coba lagi.")
		os.Exit(1)
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Save config on exit
	if err := cfg.Save(); err != nil {
		fmt.Fprintf(os.Stderr, "Gagal menyimpan konfigurasi: %v\n", err)
	}
}
