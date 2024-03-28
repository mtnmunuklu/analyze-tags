package analytics_test

import (
	"os"
	"testing"

	"github.com/mtnmunuklu/analyze-tags/analytics"
)

func TestChartGenerator_Generate(t *testing.T) {
	params := analytics.ChartParams{
		Type:   analytics.TreeChart,
		Data:   map[string][]string{"Rule1": {"tag1", "tag3", "tag2"}, "Rule2": {"tag1", "tag2"}},
		Title:  "Tree Chart Test",
		Output: "./data/output/test_tree_chart.html",
	}

	generator, err := analytics.GenerateChart(params)
	if err != nil {
		t.Errorf("Error generating chart generator: %v", err)
	}

	if err := generator.Generate(params); err != nil {
		t.Errorf("Error generating chart: %v", err)
	}

	if _, err := os.Stat(params.Output); os.IsNotExist(err) {
		t.Errorf("Chart file was not created: %v", err)
	}

}
