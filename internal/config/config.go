package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// ProviderConfig holds settings for a single AI provider.
type ProviderConfig struct {
	APIKey  string   `yaml:"api_key"`
	BaseURL string   `yaml:"base_url"`
	Model   string   `yaml:"model"`
	Models  []string `yaml:"models"`
	// Kind selects the wire protocol: "openai" (OpenAI-compatible) or "anthropic".
	Kind     string   `yaml:"kind"`
	Fallback []string `yaml:"fallback"`
}

// AgentConfig holds agent behaviour tuning.
type AgentConfig struct {
	AutoApprove   bool    `yaml:"auto_approve"`
	MaxToolCalls  int     `yaml:"max_tool_calls"`
	MaxIterations int     `yaml:"max_iterations"`
	Temperature   float64 `yaml:"temperature"`
}

// TUIConfig holds appearance options.
type TUIConfig struct {
	Theme       string `yaml:"theme"`
	Compact     bool   `yaml:"compact_mode"`
	ShowTokens  bool   `yaml:"show_token_usage"`
	Language    string `yaml:"language"`
}

// Config is the top-level configuration.
type Config struct {
	Provider  string                    `yaml:"provider"`
	Model     string                    `yaml:"model"`
	Providers map[string]ProviderConfig `yaml:"providers"`
	Agent     AgentConfig               `yaml:"agent"`
	TUI       TUIConfig                 `yaml:"tui"`

	path string `yaml:"-"`
}

// Dir returns the config directory (~/.config/termux-vibe-coding).
func Dir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ".termux-vibe-coding"
	}
	return filepath.Join(home, ".config", "termux-vibe-coding")
}

// Path returns the config file path.
func Path() string {
	return filepath.Join(Dir(), "config.yaml")
}

// Default returns a config populated with sensible defaults.
func Default() *Config {
	return &Config{
		Provider: "openrouter",
		Model:    "auto",
		Providers: map[string]ProviderConfig{
			"openrouter": {
				BaseURL: "https://openrouter.ai/api/v1",
				Kind:    "openai",
				Model:   "auto",
				Models: []string{
					"auto",
					"anthropic/claude-sonnet-4",
					"openai/gpt-4o",
					"google/gemini-2.5-pro",
					"deepseek/deepseek-chat",
				},
				Fallback: []string{"openai", "anthropic"},
			},
			"openai": {
				BaseURL: "https://api.openai.com/v1",
				Kind:    "openai",
				Model:   "gpt-4o",
				Models:  []string{"gpt-4o", "gpt-4o-mini", "o4-mini"},
			},
			"anthropic": {
				BaseURL: "https://api.anthropic.com/v1",
				Kind:    "anthropic",
				Model:   "claude-sonnet-4-20250514",
				Models:  []string{"claude-sonnet-4-20250514", "claude-opus-4-20250514"},
			},
			"gemini": {
				BaseURL: "https://generativelanguage.googleapis.com/v1beta/openai",
				Kind:    "openai",
				Model:   "gemini-2.5-pro",
				Models:  []string{"gemini-2.5-pro", "gemini-2.5-flash"},
			},
			"deepseek": {
				BaseURL: "https://api.deepseek.com/v1",
				Kind:    "openai",
				Model:   "deepseek-chat",
				Models:  []string{"deepseek-chat", "deepseek-reasoner"},
			},
			"groq": {
				BaseURL: "https://api.groq.com/openai/v1",
				Kind:    "openai",
				Model:   "llama-3.3-70b-versatile",
				Models:  []string{"llama-3.3-70b-versatile", "moonshotai/kimi-k2-instruct"},
			},
			"ollama": {
				BaseURL: "http://localhost:11434/v1",
				Kind:    "openai",
				Model:   "qwen2.5-coder",
				Models:  []string{"qwen2.5-coder", "deepseek-coder-v2", "codellama"},
			},
		},
		Agent: AgentConfig{
			AutoApprove:   false,
			MaxToolCalls:  25,
			MaxIterations: 50,
			Temperature:   0.2,
		},
		TUI: TUIConfig{
			Theme:      "catppuccin",
			Compact:    false,
			ShowTokens: true,
			Language:   "id",
		},
	}
}

// Load reads the config file, merging it over the defaults. If no file exists,
// defaults are returned (and the file is created on first Save).
func Load() (*Config, error) {
	cfg := Default()
	cfg.path = Path()

	data, err := os.ReadFile(cfg.path)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		// File doesn't exist yet — still need to resolve env vars.
		cfg.resolveEnv()
		return cfg, nil
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}

	cfg.resolveEnv()
	return cfg, nil
}

// resolveEnv expands ${VAR} references and falls back to well-known env vars.
func (c *Config) resolveEnv() {
	envFor := map[string]string{
		"openrouter": "OPENROUTER_API_KEY",
		"openai":     "OPENAI_API_KEY",
		"anthropic":  "ANTHROPIC_API_KEY",
		"gemini":     "GEMINI_API_KEY",
		"deepseek":   "DEEPSEEK_API_KEY",
		"groq":       "GROQ_API_KEY",
	}
	for name, pc := range c.Providers {
		pc.APIKey = expand(pc.APIKey)
		if pc.APIKey == "" {
			if env, ok := envFor[name]; ok {
				pc.APIKey = os.Getenv(env)
			}
		}
		c.Providers[name] = pc
	}
}

func expand(s string) string {
	if strings.HasPrefix(s, "${") && strings.HasSuffix(s, "}") {
		return os.Getenv(s[2 : len(s)-1])
	}
	return s
}

// Save writes the config to disk, creating directories as needed.
func (c *Config) Save() error {
	if c.path == "" {
		c.path = Path()
	}
	if err := os.MkdirAll(filepath.Dir(c.path), 0o755); err != nil {
		return err
	}
	data, err := yaml.Marshal(c)
	if err != nil {
		return err
	}
	return os.WriteFile(c.path, data, 0o644)
}

// Current returns the active provider config and model.
func (c *Config) Current() (ProviderConfig, string) {
	pc := c.Providers[c.Provider]
	model := c.Model
	if model == "" || model == "auto" {
		model = pc.Model
	}
	return pc, model
}

// ProviderNames returns a stable sorted list of configured provider names.
func (c *Config) ProviderNames() []string {
	order := []string{"openrouter", "openai", "anthropic", "gemini", "deepseek", "groq", "ollama"}
	var names []string
	seen := map[string]bool{}
	for _, n := range order {
		if _, ok := c.Providers[n]; ok {
			names = append(names, n)
			seen[n] = true
		}
	}
	for n := range c.Providers {
		if !seen[n] {
			names = append(names, n)
		}
	}
	return names
}
