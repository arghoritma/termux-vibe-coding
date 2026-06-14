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

// OpenAIProvider talks to any OpenAI-compatible /chat/completions endpoint.
// This covers OpenRouter, OpenAI, Gemini (openai-compat), DeepSeek, Groq, Ollama.
type OpenAIProvider struct {
	name    string
	baseURL string
	apiKey  string
	client  *http.Client
}

// NewOpenAI creates an OpenAI-compatible provider.
func NewOpenAI(name, baseURL, apiKey string) *OpenAIProvider {
	return &OpenAIProvider{
		name:    name,
		baseURL: strings.TrimRight(baseURL, "/"),
		apiKey:  apiKey,
		client:  &http.Client{Timeout: 180 * time.Second},
	}
}

func (p *OpenAIProvider) Name() string { return p.name }

type oaiMessage struct {
	Role       string         `json:"role"`
	Content    any            `json:"content,omitempty"`
	ToolCalls  []oaiToolCall  `json:"tool_calls,omitempty"`
	ToolCallID string         `json:"tool_call_id,omitempty"`
	Name       string         `json:"name,omitempty"`
}

type oaiToolCall struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Function struct {
		Name      string `json:"name"`
		Arguments string `json:"arguments"`
	} `json:"function"`
}

type oaiTool struct {
	Type     string `json:"type"`
	Function struct {
		Name        string         `json:"name"`
		Description string         `json:"description"`
		Parameters  map[string]any `json:"parameters"`
	} `json:"function"`
}

type oaiRequest struct {
	Model       string       `json:"model"`
	Messages    []oaiMessage `json:"messages"`
	Tools       []oaiTool    `json:"tools,omitempty"`
	Temperature float64      `json:"temperature"`
}

type oaiResponse struct {
	Choices []struct {
		Message oaiMessage `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Model string `json:"model"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error"`
}

func (p *OpenAIProvider) Complete(ctx context.Context, req Request) (*Response, error) {
	body := oaiRequest{
		Model:       req.Model,
		Temperature: req.Temperature,
	}

	for _, m := range req.Messages {
		om := oaiMessage{Role: m.Role}
		if m.Content != "" {
			om.Content = m.Content
		}
		if m.Role == RoleTool {
			om.ToolCallID = m.ToolCallID
			om.Name = m.Name
		}
		for _, tc := range m.ToolCalls {
			otc := oaiToolCall{ID: tc.ID, Type: "function"}
			otc.Function.Name = tc.Name
			otc.Function.Arguments = tc.Arguments
			om.ToolCalls = append(om.ToolCalls, otc)
		}
		body.Messages = append(body.Messages, om)
	}

	for _, t := range req.Tools {
		ot := oaiTool{Type: "function"}
		ot.Function.Name = t.Name
		ot.Function.Description = t.Description
		ot.Function.Parameters = t.Parameters
		body.Tools = append(body.Tools, ot)
	}

	raw, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, p.baseURL+"/chat/completions", bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", "application/json")
	if p.apiKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+p.apiKey)
	}
	// OpenRouter ranking headers (harmless elsewhere).
	httpReq.Header.Set("HTTP-Referer", "https://github.com/rozaq/termux-vibe-coding")
	httpReq.Header.Set("X-Title", "termux-vibe-coding")

	resp, err := p.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request gagal: %w", err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("provider %s error %d: %s", p.name, resp.StatusCode, snippet(data))
	}

	var out oaiResponse
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, fmt.Errorf("parsing response: %w (%s)", err, snippet(data))
	}
	if out.Error != nil {
		return nil, fmt.Errorf("provider %s: %s", p.name, out.Error.Message)
	}
	if len(out.Choices) == 0 {
		return nil, fmt.Errorf("provider %s: tidak ada jawaban", p.name)
	}

	msg := out.Choices[0].Message
	r := &Response{
		Model: out.Model,
		Usage: Usage{
			PromptTokens:     out.Usage.PromptTokens,
			CompletionTokens: out.Usage.CompletionTokens,
			TotalTokens:      out.Usage.TotalTokens,
		},
	}
	if s, ok := msg.Content.(string); ok {
		r.Content = s
	}
	for _, tc := range msg.ToolCalls {
		r.ToolCalls = append(r.ToolCalls, ToolCall{
			ID:        tc.ID,
			Name:      tc.Function.Name,
			Arguments: tc.Function.Arguments,
		})
	}
	return r, nil
}

func snippet(b []byte) string {
	s := string(b)
	if len(s) > 400 {
		return s[:400] + "..."
	}
	return s
}
