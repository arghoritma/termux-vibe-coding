package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// AnthropicProvider talks to the native Anthropic Messages API.
type AnthropicProvider struct {
	name    string
	baseURL string
	apiKey  string
	client  *http.Client
}

// NewAnthropic creates a native Anthropic provider.
func NewAnthropic(name, baseURL, apiKey string) *AnthropicProvider {
	return &AnthropicProvider{
		name:    name,
		baseURL: strings.TrimRight(baseURL, "/"),
		apiKey:  apiKey,
		client:  &http.Client{Timeout: 180 * time.Second},
	}
}

func (p *AnthropicProvider) Name() string { return p.name }

type antContent struct {
	Type      string          `json:"type"`
	Text      string          `json:"text,omitempty"`
	ID        string          `json:"id,omitempty"`
	Name      string          `json:"name,omitempty"`
	Input     json.RawMessage `json:"input,omitempty"`
	ToolUseID string          `json:"tool_use_id,omitempty"`
	Content   string          `json:"content,omitempty"`
}

type antMessage struct {
	Role    string       `json:"role"`
	Content []antContent `json:"content"`
}

type antTool struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	InputSchema map[string]any `json:"input_schema"`
}

type antRequest struct {
	Model       string       `json:"model"`
	MaxTokens   int          `json:"max_tokens"`
	System      string       `json:"system,omitempty"`
	Messages    []antMessage `json:"messages"`
	Tools       []antTool    `json:"tools,omitempty"`
	Temperature float64      `json:"temperature"`
}

type antResponse struct {
	Content []antContent `json:"content"`
	Model   string       `json:"model"`
	Usage   struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

func (p *AnthropicProvider) Complete(ctx context.Context, req Request) (*Response, error) {
	body := antRequest{
		Model:       req.Model,
		MaxTokens:   8192,
		Temperature: req.Temperature,
	}

	for _, m := range req.Messages {
		switch m.Role {
		case RoleSystem:
			if body.System != "" {
				body.System += "\n\n"
			}
			body.System += m.Content
		case RoleUser:
			body.Messages = append(body.Messages, antMessage{
				Role:    "user",
				Content: []antContent{{Type: "text", Text: m.Content}},
			})
		case RoleAssistant:
			var blocks []antContent
			if m.Content != "" {
				blocks = append(blocks, antContent{Type: "text", Text: m.Content})
			}
			for _, tc := range m.ToolCalls {
				blocks = append(blocks, antContent{
					Type:  "tool_use",
					ID:    tc.ID,
					Name:  tc.Name,
					Input: json.RawMessage(orEmptyJSON(tc.Arguments)),
				})
			}
			body.Messages = append(body.Messages, antMessage{Role: "assistant", Content: blocks})
		case RoleTool:
			body.Messages = append(body.Messages, antMessage{
				Role: "user",
				Content: []antContent{{
					Type:      "tool_result",
					ToolUseID: m.ToolCallID,
					Content:   m.Content,
				}},
			})
		}
	}

	for _, t := range req.Tools {
		body.Tools = append(body.Tools, antTool{
			Name:        t.Name,
			Description: t.Description,
			InputSchema: t.Parameters,
		})
	}

	raw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/messages", bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("x-api-key", p.apiKey)
	httpReq.Header.Set("anthropic-version", "2023-06-01")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request gagal: %w", err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("anthropic error %d: %s", resp.StatusCode, snippet(data))
	}

	var out antResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("parsing response: %w (%s)", err, snippet(data))
	}
	if out.Error != nil {
		return nil, fmt.Errorf("anthropic: %s", out.Error.Message)
	}

	r := &Response{
		Model: out.Model,
		Usage: Usage{
			PromptTokens:     out.Usage.InputTokens,
			CompletionTokens: out.Usage.OutputTokens,
			TotalTokens:      out.Usage.InputTokens + out.Usage.OutputTokens,
		},
	}
	for _, c := range out.Content {
		switch c.Type {
		case "text":
			r.Content += c.Text
		case "tool_use":
			r.ToolCalls = append(r.ToolCalls, ToolCall{
				ID:        c.ID,
				Name:      c.Name,
				Arguments: string(c.Input),
			})
		}
	}
	return r, nil
}

func orEmptyJSON(s string) string {
	if strings.TrimSpace(s) == "" {
		return "{}"
	}
	return s
}
