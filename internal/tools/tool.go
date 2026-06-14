package tools

import (
	"context"

	"github.com/rozaq/termux-vibe-coding/internal/ai"
)

// Tool is an action the agent can perform.
type Tool interface {
	Name() string
	Description() string
	// Schema returns a JSON Schema object describing the arguments.
	Schema() map[string]any
	// Dangerous reports whether the tool needs user approval before running.
	Dangerous() bool
	// Execute runs the tool with raw JSON arguments and returns a result string.
	Execute(ctx context.Context, args string) (string, error)
}

// Registry holds the available tools.
type Registry struct {
	tools map[string]Tool
	order []string
}

// NewRegistry builds the default tool set rooted at workdir.
func NewRegistry(workdir string) *Registry {
	r := &Registry{tools: map[string]Tool{}}
	r.register(&ReadTool{root: workdir})
	r.register(&WriteTool{root: workdir})
	r.register(&EditTool{root: workdir})
	r.register(&GlobTool{root: workdir})
	r.register(&GrepTool{root: workdir})
	r.register(&BashTool{root: workdir})
	return r
}

func (r *Registry) register(t Tool) {
	r.tools[t.Name()] = t
	r.order = append(r.order, t.Name())
}

// Get returns a tool by name.
func (r *Registry) Get(name string) (Tool, bool) {
	t, ok := r.tools[name]
	return t, ok
}

// Defs returns the AI tool definitions for all registered tools.
func (r *Registry) Defs() []ai.ToolDef {
	var defs []ai.ToolDef
	for _, name := range r.order {
		t := r.tools[name]
		defs = append(defs, ai.ToolDef{
			Name:        t.Name(),
			Description: t.Description(),
			Parameters:  t.Schema(),
		})
	}
	return defs
}

// helper to build a JSON schema object.
func object(props map[string]any, required ...string) map[string]any {
	if required == nil {
		required = []string{}
	}
	return map[string]any{
		"type":       "object",
		"properties": props,
		"required":   required,
	}
}

func strProp(desc string) map[string]any {
	return map[string]any{"type": "string", "description": desc}
}
