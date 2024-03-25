package analytics_test

import (
	"os"
	"testing"

	"github.com/mtnmunuklu/analyze-tags/analytics"
)

func TestChartGenerator_Generate(t *testing.T) {
	// Mock data
	params := analytics.ChartParams{
		Type:   analytics.GraphChart,
		Data:   map[string][]string{"Rule1": {"tag1", "tag3"}, "Rule2": {"tag1", "tag2"}},
		Title:  "Garph Chart Test",
		Output: "./data/output/test_graph_chart.html",
	}

	// Generate chart generator
	generator, err := analytics.GenerateChart(params)
	if err != nil {
		t.Errorf("Error generating chart generator: %v", err)
	}

	// Generate chart
	if err := generator.Generate(params); err != nil {
		t.Errorf("Error generating chart: %v", err)
	}

	// Check if the chart file was created
	if _, err := os.Stat(params.Output); os.IsNotExist(err) {
		t.Errorf("Chart file was not created: %v", err)
	}

	// Clean up: delete the created chart file
	//if err := os.Remove(params.Output); err != nil {
	//	t.Errorf("Error cleaning up: %v", err)
	//}
}
