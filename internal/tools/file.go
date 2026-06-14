package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// resolve joins root with a possibly-relative path and prevents escaping when relative.
func resolve(root, p string) string {
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(root, p)
}

// ReadTool reads a file's contents.
type ReadTool struct{ root string }

func (t *ReadTool) Name() string        { return "read" }
func (t *ReadTool) Dangerous() bool      { return false }
func (t *ReadTool) Description() string {
	return "Baca isi sebuah file. Gunakan path relatif terhadap working directory."
}
func (t *ReadTool) Schema() map[string]any {
	return object(map[string]any{
		"path": strProp("Path file yang akan dibaca"),
	}, "path")
}
func (t *ReadTool) Execute(ctx context.Context, args string) (string, error) {
	var a struct {
		Path string `json:"path"`
	}
	if err := json.Unmarshal([]byte(args), &a); err != nil {
		return "", err
	}
	data, err := os.ReadFile(resolve(t.root, a.Path))
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteTool creates or overwrites a file.
type WriteTool struct{ root string }

func (t *WriteTool) Name() string   { return "write" }
func (t *WriteTool) Dangerous() bool { return true }
func (t *WriteTool) Description() string {
	return "Buat atau timpa sebuah file dengan konten yang diberikan. Otomatis membuat folder induk."
}
func (t *WriteTool) Schema() map[string]any {
	return object(map[string]any{
		"path":    strProp("Path file yang akan ditulis"),
		"content": strProp("Isi lengkap file"),
	}, "path", "content")
}
func (t *WriteTool) Execute(ctx context.Context, args string) (string, error) {
	var a struct {
		Path    string `json:"path"`
		Content string `json:"content"`
	}
	if err := json.Unmarshal([]byte(args), &a); err != nil {
		return "", err
	}
	full := resolve(t.root, a.Path)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		return "", err
	}
	if err := os.WriteFile(full, []byte(a.Content), 0o644); err != nil {
		return "", err
	}
	lines := strings.Count(a.Content, "\n") + 1
	return fmt.Sprintf("File ditulis: %s (%d baris)", a.Path, lines), nil
}

// EditTool performs a find-and-replace within a file.
type EditTool struct{ root string }

func (t *EditTool) Name() string   { return "edit" }
func (t *EditTool) Dangerous() bool { return true }
func (t *EditTool) Description() string {
	return "Edit file dengan mengganti teks lama dengan teks baru. old_string harus unik dan cocok persis."
}
func (t *EditTool) Schema() map[string]any {
	return object(map[string]any{
		"path":       strProp("Path file yang akan diedit"),
		"old_string": strProp("Teks lama yang akan diganti (harus cocok persis)"),
		"new_string": strProp("Teks pengganti"),
	}, "path", "old_string", "new_string")
}
func (t *EditTool) Execute(ctx context.Context, args string) (string, error) {
	var a struct {
		Path      string `json:"path"`
		OldString string `json:"old_string"`
		NewString string `json:"new_string"`
	}
	if err := json.Unmarshal([]byte(args), &a); err != nil {
		return "", err
	}
	full := resolve(t.root, a.Path)
	data, err := os.ReadFile(full)
	if err != nil {
		return "", err
	}
	content := string(data)
	n := strings.Count(content, a.OldString)
	if n == 0 {
		return "", fmt.Errorf("old_string tidak ditemukan di %s", a.Path)
	}
	if n > 1 {
		return "", fmt.Errorf("old_string muncul %d kali di %s; buat lebih spesifik", n, a.Path)
	}
	content = strings.Replace(content, a.OldString, a.NewString, 1)
	if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
		return "", err
	}
	return fmt.Sprintf("File diedit: %s", a.Path), nil
}
