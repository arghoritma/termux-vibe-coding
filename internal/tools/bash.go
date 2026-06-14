package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

// BashTool runs a shell command in the working directory.
type BashTool struct{ root string }

func (t *BashTool) Name() string   { return "bash" }
func (t *BashTool) Dangerous() bool { return true }
func (t *BashTool) Description() string {
	return "Jalankan perintah shell di working directory (npm, git, ls, dll). Output digabung stdout+stderr."
}
func (t *BashTool) Schema() map[string]any {
	return object(map[string]any{
		"command": strProp("Perintah shell yang dijalankan"),
	}, "command")
}
func (t *BashTool) Execute(ctx context.Context, args string) (string, error) {
	var a struct {
		Command string `json:"command"`
	}
	if err := json.Unmarshal([]byte(args), &a); err != nil {
		return "", err
	}

	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "sh", "-c", a.Command)
	cmd.Dir = t.root
	out, err := cmd.CombinedOutput()
	result := string(out)
	if err != nil {
		return result, fmt.Errorf("perintah selesai dengan error: %v", err)
	}
	if result == "" {
		result = "(tidak ada output)"
	}
	return result, nil
}
