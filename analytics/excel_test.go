package analytics_test

import (
	"os"
	"testing"

	"github.com/mtnmunuklu/analyze-tags/analytics"
	"github.com/stretchr/testify/assert"
)

func TestToExcel(t *testing.T) {
	data := map[string][]string{
		"Rule1": {"tag1", "tag3", "tag2"},
		"Rule2": {"tag1", "tag2"},
	}

	params := analytics.ExcelParams{
		SheetName: "TestSheet",
		Data:      data,
		Output:    "./data/output/test/output.xlsx",
	}

	err := params.ToExcel()
	defer func() {
		err := os.Remove("./data/output/test/output.xlsx")
		if err != nil {
			t.Errorf("Error deleting test file: %v", err)
		}
	}()

	assert.Nil(t, err, "Expected error to be nil")
}
