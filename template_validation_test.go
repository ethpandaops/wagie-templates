package wagietemplates_test

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ethpandaops/wagie"
)

func TestTemplatesValidateAgainstCombinedLibrary(t *testing.T) {
	t.Parallel()

	coreRoot := os.Getenv("WAGIE_CORE_DIR")
	if coreRoot == "" {
		coreRoot = filepath.Join("..", "wagie")
	}

	coreTemplatesDir := filepath.Clean(filepath.Join(coreRoot, "templates"))
	if _, err := os.Stat(coreTemplatesDir); err != nil {
		t.Fatalf("core templates dir %s not available: %v", coreTemplatesDir, err)
	}

	dirs := []struct {
		path   string
		source string
	}{
		{path: coreTemplatesDir, source: "wagie-core"},
		{path: "ethereum", source: "wagie-templates"},
		{path: "code", source: "wagie-templates"},
		{path: "research", source: "wagie-templates"},
	}

	files := make([]wagie.TemplateFile, 0, 64)
	for _, dir := range dirs {
		loaded, err := wagie.LoadTemplateFilesRecursive(dir.path)
		if err != nil {
			t.Fatalf("load %s: %v", dir.path, err)
		}

		for _, file := range loaded {
			file.Source = dir.source
			files = append(files, file)
		}
	}

	results, err := wagie.ValidateTemplateFiles(
		context.Background(),
		slog.New(slog.DiscardHandler),
		files,
	)
	if err != nil {
		t.Fatalf("validate files: %v", err)
	}

	var failures []string
	for _, result := range results {
		if result.Valid {
			continue
		}

		lines := make([]string, 0, len(result.Errors)+1)
		header := result.Path
		if result.Name != "" {
			header = fmt.Sprintf("%s (%s@%s)", result.Path, result.Name, result.Version)
		}
		lines = append(lines, header)
		for _, err := range result.Errors {
			if err.Line > 0 {
				lines = append(lines, fmt.Sprintf("  - [%s] line %d: %s", err.Type, err.Line, err.Message))
				continue
			}
			lines = append(lines, fmt.Sprintf("  - [%s] %s", err.Type, err.Message))
		}
		failures = append(failures, strings.Join(lines, "\n"))
	}

	if len(failures) > 0 {
		t.Fatalf("template validation failures:\n%s", strings.Join(failures, "\n\n"))
	}
}
