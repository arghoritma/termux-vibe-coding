package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// GlobTool finds files matching a glob pattern.
type GlobTool struct{ root string }

func (t *GlobTool) Name() string   { return "glob" }
func (t *GlobTool) Dangerous() bool { return false }
func (t *GlobTool) Description() string {
	return "Cari file berdasarkan pola glob, mis. '**/*.go' atau 'src/*.js'."
}
func (t *GlobTool) Schema() map[string]any {
	return object(map[string]any{
		"pattern": strProp("Pola glob, mis. **/*.js"),
	}, "pattern")
}
func (t *GlobTool) Execute(ctx context.Context, args string) (string, error) {
	var a struct {
		Pattern string `json:"pattern"`
	}
	if err := json.Unmarshal([]byte(args), &a); err != nil {
		return "", err
	}

	// Support "**" by walking the tree and matching the base pattern.
	doublestar := strings.Contains(a.Pattern, "**")
	base := strings.ReplaceAll(a.Pattern, "**/", "")
	base = strings.ReplaceAll(base, "**", "")

	var matches []string
	err := filepath.WalkDir(t.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if name == ".git" || name == "node_modules" || name == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}
		rel, _ := filepath.Rel(t.root, path)
		var ok bool
		if doublestar {
			ok, _ = filepath.Match(base, filepath.Base(path))
		} else {
			ok, _ = filepath.Match(a.Pattern, rel)
			if !ok {
				ok, _ = filepath.Match(a.Pattern, filepath.Base(path))
			}
		}
		if ok {
			matches = append(matches, rel)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if len(matches) == 0 {
		return "Tidak ada file yang cocok.", nil
	}
	if len(matches) > 200 {
		matches = matches[:200]
	}
	return strings.Join(matches, "\n"), nil
}

// GrepTool searches file contents with a regular expression.
type GrepTool struct{ root string }

func (t *GrepTool) Name() string   { return "grep" }
func (t *GrepTool) Dangerous() bool { return false }
func (t *GrepTool) Description() string {
	return "Cari teks dalam file menggunakan regular expression. Mengembalikan path:baris:isi."
}
func (t *GrepTool) Schema() map[string]any {
	return object(map[string]any{
		"pattern": strProp("Regex yang dicari"),
		"include": strProp("Opsional: filter glob nama file, mis. *.go"),
	}, "pattern")
}
func (t *GrepTool) Execute(ctx context.Context, args string) (string, error) {
	var a struct {
		Pattern string `json:"pattern"`
		Include string `json:"include"`
	}
	if err := json.Unmarshal([]byte(args), &a); err != nil {
		return "", err
	}
	re, err := regexp.Compile(a.Pattern)
	if err != nil {
		return "", fmt.Errorf("regex tidak valid: %w", err)
	}

	var results []string
	count := 0
	err = filepath.WalkDir(t.root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if name == ".git" || name == "node_modules" || name == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}
		if count >= 200 {
			return filepath.SkipAll
		}
		if a.Include != "" {
			ok, _ := filepath.Match(a.Include, filepath.Base(path))
			if !ok {
				return nil
			}
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return nil
		}
		rel, _ := filepath.Rel(t.root, path)
		for i, line := range strings.Split(string(data), "\n") {
			if re.MatchString(line) {
				trimmed := strings.TrimSpace(line)
				if len(trimmed) > 200 {
					trimmed = trimmed[:200]
				}
				results = append(results, fmt.Sprintf("%s:%d: %s", rel, i+1, trimmed))
				count++
				if count >= 200 {
					return filepath.SkipAll
				}
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if len(results) == 0 {
		return "Tidak ada yang cocok.", nil
	}
	return strings.Join(results, "\n"), nil
}
