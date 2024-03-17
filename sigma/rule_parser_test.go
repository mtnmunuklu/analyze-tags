package sigma

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseRule(t *testing.T) {
	err := filepath.Walk("./data/rules/", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".rule.yml") {
			return nil
		}

		t.Run(strings.TrimSuffix(filepath.Base(path), ".rule.yml"), func(t *testing.T) {
			contents, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("failed reading test input: %v", err)
			}

			_, err = ParseRule(contents)
			if err != nil {
				t.Fatalf("error parsing rule: %v", err)
			}

		})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
