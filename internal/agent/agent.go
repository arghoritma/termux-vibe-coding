package agent

import (
	"context"
	"fmt"

	"github.com/arghoritma/termux-vibe-coding/internal/ai"
	"github.com/arghoritma/termux-vibe-coding/internal/config"
	"github.com/arghoritma/termux-vibe-coding/internal/tools"
)

// Handler receives callbacks during a turn so the UI can render progress.
type Handler struct {
	// OnText is called with assistant prose output.
	OnText func(text string)
	// OnTool is called when the agent decides to invoke a tool.
	OnTool func(name, args string)
	// OnResult is called after a tool finishes.
	OnResult func(name, result string, err error)
	// Approve is called for dangerous tools; return true to allow.
	Approve func(name, args string) bool
}

// Agent drives the think -> act -> observe loop.
type Agent struct {
	provider ai.Provider
	tools    *tools.Registry
	cfg      *config.Config
	model    string
	workdir  string

	messages   []ai.Message
	TotalUsage ai.Usage
}

// New creates an agent for the given provider and model.
func New(cfg *config.Config, provider ai.Provider, model, workdir string) *Agent {
	a := &Agent{
		provider: provider,
		tools:    tools.NewRegistry(workdir),
		cfg:      cfg,
		model:    model,
		workdir:  workdir,
	}
	a.messages = []ai.Message{{
		Role:    ai.RoleSystem,
		Content: systemPrompt(workdir, cfg.TUI.Language),
	}}
	return a
}

// SetProvider swaps the active provider/model, preserving conversation history.
func (a *Agent) SetProvider(provider ai.Provider, model string) {
	a.provider = provider
	a.model = model
}

// Model returns the active model name.
func (a *Agent) Model() string { return a.model }

// Provider returns the active provider name.
func (a *Agent) Provider() string { return a.provider.Name() }

// Send processes one user message, looping through tool calls until the model
// produces a final text answer (or limits are hit).
func (a *Agent) Send(ctx context.Context, userText string, h Handler) error {
	a.messages = append(a.messages, ai.Message{Role: ai.RoleUser, Content: userText})

	for iter := 0; iter < a.cfg.Agent.MaxIterations; iter++ {
		resp, err := a.provider.Complete(ctx, ai.Request{
			Model:       a.model,
			Messages:    a.messages,
			Tools:       a.tools.Defs(),
			Temperature: a.cfg.Agent.Temperature,
		})
		if err != nil {
			return err
		}

		a.TotalUsage.TotalTokens += resp.Usage.TotalTokens
		a.TotalUsage.PromptTokens += resp.Usage.PromptTokens
		a.TotalUsage.CompletionTokens += resp.Usage.CompletionTokens

		// Record the assistant turn.
		a.messages = append(a.messages, ai.Message{
			Role:      ai.RoleAssistant,
			Content:   resp.Content,
			ToolCalls: resp.ToolCalls,
		})

		if resp.Content != "" && h.OnText != nil {
			h.OnText(resp.Content)
		}

		// No tool calls -> turn is complete.
		if len(resp.ToolCalls) == 0 {
			return nil
		}

		// Execute each requested tool.
		for _, tc := range resp.ToolCalls {
			result := a.runTool(ctx, tc, h)
			a.messages = append(a.messages, ai.Message{
				Role:       ai.RoleTool,
				Content:    result,
				ToolCallID: tc.ID,
				Name:       tc.Name,
			})
		}
	}
	return fmt.Errorf("batas iterasi tercapai (%d)", a.cfg.Agent.MaxIterations)
}

func (a *Agent) runTool(ctx context.Context, tc ai.ToolCall, h Handler) string {
	tool, ok := a.tools.Get(tc.Name)
	if !ok {
		if h.OnResult != nil {
			h.OnResult(tc.Name, "", fmt.Errorf("tool tidak dikenal"))
		}
		return "Error: tool tidak dikenal: " + tc.Name
	}

	if h.OnTool != nil {
		h.OnTool(tc.Name, tc.Arguments)
	}

	if tool.Dangerous() && !a.cfg.Agent.AutoApprove {
		if h.Approve != nil && !h.Approve(tc.Name, tc.Arguments) {
			msg := "Dibatalkan oleh pengguna."
			if h.OnResult != nil {
				h.OnResult(tc.Name, msg, nil)
			}
			return msg
		}
	}

	out, err := tool.Execute(ctx, tc.Arguments)
	if h.OnResult != nil {
		h.OnResult(tc.Name, out, err)
	}
	if err != nil {
		return fmt.Sprintf("Error: %v\n%s", err, out)
	}
	return out
}
