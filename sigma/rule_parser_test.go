package sigma_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mtnmunuklu/analyze-tags/sigma"
)

func TestParseRule(t *testing.T) {
	err := filepath.Walk("./data/rules/", func(path string, info os.FileInfo, err error) error {
		if !strings.HasSuffix(path, ".yml") {
			return nil
		}

		t.Run(strings.TrimSuffix(filepath.Base(path), ".yml"), func(t *testing.T) {
			contents, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("failed reading test input: %v", err)
			}

			_, err = sigma.ParseRule(contents)
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
