package ai

import (
	"fmt"

	"github.com/rozaq/termux-vibe-coding/internal/config"
)

// Build creates a Provider for the named provider in the config.
func Build(cfg *config.Config, name string) (Provider, error) {
	pc, ok := cfg.Providers[name]
	if !ok {
		return nil, fmt.Errorf("provider %q tidak dikenal", name)
	}
	switch pc.Kind {
	case "anthropic":
		if pc.APIKey == "" {
			return nil, fmt.Errorf("API key untuk %s belum diset (export ANTHROPIC_API_KEY)", name)
		}
		return NewAnthropic(name, pc.BaseURL, pc.APIKey), nil
	default: // openai-compatible
		// Ollama runs locally and needs no key.
		if pc.APIKey == "" && name != "ollama" {
			return nil, fmt.Errorf("API key untuk %s belum diset", name)
		}
		return NewOpenAI(name, pc.BaseURL, pc.APIKey), nil
	}
}
