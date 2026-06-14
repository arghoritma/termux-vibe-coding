package ai

import "context"

// Role constants for chat messages.
const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistant = "assistant"
	RoleTool      = "tool"
)

// ToolDef describes a tool/function the model may call.
type ToolDef struct {
	Name        string
	Description string
	// Parameters is a JSON Schema object describing the arguments.
	Parameters map[string]any
}

// ToolCall is a request from the model to invoke a tool.
type ToolCall struct {
	ID        string
	Name      string
	Arguments string // raw JSON
}

// Message is a single chat message.
type Message struct {
	Role       string
	Content    string
	ToolCalls  []ToolCall // assistant -> tool calls
	ToolCallID string     // for role=tool, which call this answers
	Name       string     // tool name for role=tool
}

// Usage reports token consumption.
type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

// Response is the result of a completion.
type Response struct {
	Content   string
	ToolCalls []ToolCall
	Usage     Usage
	Model     string
}

// Request bundles inputs for a completion.
type Request struct {
	Model       string
	Messages    []Message
	Tools       []ToolDef
	Temperature float64
}

// Provider is the interface every AI backend implements.
type Provider interface {
	Name() string
	Complete(ctx context.Context, req Request) (*Response, error)
}
