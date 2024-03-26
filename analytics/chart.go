package analytics

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type ChartType string

const (
	BarChart       ChartType = "bar"
	LineChart      ChartType = "line"
	ScatterPlot    ChartType = "scatter"
	PieChart       ChartType = "pie"
	BoxPlotChart   ChartType = "boxplot"
	HeatmapChart   ChartType = "heatmap"
	RadarChart     ChartType = "radar"
	FunnelChart    ChartType = "funnel"
	WordCloudChart ChartType = "wordcloud"
	TreemapChart   ChartType = "treemap"
	GraphChart     ChartType = "graph"
	TreeChart      ChartType = "tree"
)

type ChartParams struct {
	Type   ChartType
	Data   map[string][]string
	Title  string
	Output string
}

func renderChartToFile(chart components.Charter, outputPath string) error {
	f, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer f.Close()

	page := components.NewPage()
	page.AddCharts(chart)
	page.Render(f)

	return nil
}

func GenerateChart(params ChartParams) (ChartGenerator, error) {
	switch params.Type {
	case BarChart:
		return &BarChartGenerator{}, nil
	case LineChart:
		return &LineChartGenerator{}, nil
	case ScatterPlot:
		return &ScatterPlotGenerator{}, nil
	case PieChart:
		return &PieChartGenerator{}, nil
	case BoxPlotChart:
		return &BoxPlotGenerator{}, nil
	case HeatmapChart:
		return &HeatmapGenerator{}, nil
	case RadarChart:
		return &RadarChartGenerator{}, nil
	case FunnelChart:
		return &FunnelChartGenerator{}, nil
	case WordCloudChart:
		return &WordCloudChartGenerator{}, nil
	case TreemapChart:
		return &TreemapChartGenerator{}, nil
	case GraphChart:
		return &GraphChartGenerator{}, nil
	case TreeChart:
		return &TreeChartGenerator{}, nil
	default:
		return nil, fmt.Errorf("unsupported chart type: %s", params.Type)
	}
}

type ChartGenerator interface {
	Generate(params ChartParams) error
}

type BarChartGenerator struct{}

func (bcg *BarChartGenerator) Generate(params ChartParams) error {

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	var xAxisData []string
	var seriesData []opts.BarData
	for tag, count := range tagCounts {
		xAxisData = append(xAxisData, tag)
		seriesData = append(seriesData, opts.BarData{Value: count})
	}

	bar.SetXAxis(xAxisData).
		AddSeries("Count", seriesData,
			charts.WithLabelOpts(opts.Label{
				Show:      true,
				Position:  "top",
				Formatter: "{c}",
			}))

	return renderChartToFile(bar, params.Output)
}

type LineChartGenerator struct{}

func (lcg *LineChartGenerator) Generate(params ChartParams) error {

	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	var xAxisData []string
	var lineData []opts.LineData
	for tag, count := range tagCounts {
		xAxisData = append(xAxisData, tag)
		lineData = append(lineData, opts.LineData{Value: count})
	}

	line.SetXAxis(xAxisData).
		AddSeries("Count", lineData)

	return renderChartToFile(line, params.Output)
}

type ScatterPlotGenerator struct{}

func (spg *ScatterPlotGenerator) Generate(params ChartParams) error {

	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	var scatterData []opts.ScatterData
	var xAxisData []string
	for tag, count := range tagCounts {
		xAxisData = append(xAxisData, tag)
		scatterData = append(scatterData, opts.ScatterData{Value: []interface{}{tag, count}})
	}

	scatter.SetXAxis(xAxisData).AddSeries("Count", scatterData)

	return renderChartToFile(scatter, params.Output)
}

type PieChartGenerator struct{}

func (pcg *PieChartGenerator) Generate(params ChartParams) error {

	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	tagCounts := make(map[string]int)
	totalCount := 0
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
			totalCount++
		}
	}

	var pieData []opts.PieData
	for tag, count := range tagCounts {

		percentage := float64(count) / float64(totalCount) * 100
		label := fmt.Sprintf("{b}: %d (%.2f%%)", count, percentage)
		pieData = append(pieData, opts.PieData{Name: tag, Value: count, Label: &opts.Label{Show: true, Formatter: label}})
	}

	pie.AddSeries("Count", pieData)

	return renderChartToFile(pie, params.Output)
}

type BoxPlotGenerator struct{}

func (bpg *BoxPlotGenerator) Generate(params ChartParams) error {

	boxplot := charts.NewBoxPlot()
	boxplot.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	var xAxisData []string
	var boxplotData []opts.BoxPlotData
	for tag, count := range tagCounts {
		xAxisData = append(xAxisData, tag)
		boxplotData = append(boxplotData, opts.BoxPlotData{Name: tag, Value: count})
	}

	boxplot.SetXAxis(xAxisData)
	boxplot.AddSeries("Count", boxplotData)

	return renderChartToFile(boxplot, params.Output)
}

type HeatmapGenerator struct{}

func (hmg *HeatmapGenerator) Generate(params ChartParams) error {

	heatmap := charts.NewHeatMap()

	heatmap.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Type:      "category",
			SplitArea: &opts.SplitArea{Show: true},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type:      "category",
			SplitArea: &opts.SplitArea{Show: true},
		}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Min:        0,
			Max:        10,
			InRange: &opts.VisualMapInRange{
				Color: []string{"#50a3ba", "#eac736", "#d94e5d"},
			},
		}),
	)

	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	var data []opts.HeatMapData
	for rule, tags := range params.Data {
		for _, tag := range tags {

			count := tagCounts[tag]
			data = append(data, opts.HeatMapData{Value: [3]interface{}{tag, rule, count}})
		}
	}

	heatmap.AddSeries("Count", data)

	return renderChartToFile(heatmap, params.Output)
}

type RadarChartGenerator struct{}

func (rcg *RadarChartGenerator) Generate(params ChartParams) error {

	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	indicators := make([]*opts.Indicator, 0)
	for rule := range params.Data {
		indicators = append(indicators, &opts.Indicator{Name: rule})
	}

	seriesData := make([]opts.RadarData, 0)
	for _, tags := range params.Data {
		var rowData []interface{}
		for _, tag := range tags {
			rowData = append(rowData, tag)
		}
		seriesData = append(seriesData, opts.RadarData{Value: rowData})
	}

	radar.SetGlobalOptions(charts.WithRadarComponentOpts(opts.RadarComponent{
		Indicator: indicators,
	}))

	radar.AddSeries("Data", seriesData)

	return renderChartToFile(radar, params.Output)
}

type FunnelChartGenerator struct{}

func (fcg *FunnelChartGenerator) Generate(params ChartParams) error {

	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	var data []opts.FunnelData
	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}
	for tag, count := range tagCounts {
		label := fmt.Sprintf("%s: %d", tag, count)
		data = append(data, opts.FunnelData{Name: label, Value: float32(count)})
	}

	funnel.AddSeries("Data", data)

	return renderChartToFile(funnel, params.Output)
}

type WordCloudChartGenerator struct{}

func (wcg *WordCloudChartGenerator) Generate(params ChartParams) error {

	wordCloud := charts.NewWordCloud()
	wordCloud.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	tagCounts := make(map[string]int)
	ruleTagCounts := make(map[string]int)
	for rule, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
			ruleTagCounts[rule]++
		}
	}

	var data []opts.WordCloudData
	processedTags := make(map[string]bool)
	for rule, tags := range params.Data {
		ruleWeight := float32(ruleTagCounts[rule])
		for _, tag := range tags {

			if _, ok := processedTags[tag]; ok {
				continue
			}
			tagWeight := float32(tagCounts[tag])

			data = append(data, opts.WordCloudData{Name: tag, Value: tagWeight})
			processedTags[tag] = true
		}

		data = append(data, opts.WordCloudData{Name: rule, Value: ruleWeight})
	}

	wordCloud.AddSeries("Data", data)

	return renderChartToFile(wordCloud, params.Output)
}

type TreemapChartGenerator struct{}

func (tcg *TreemapChartGenerator) Generate(params ChartParams) error {

	treemap := charts.NewTreeMap()
	treemap.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	var treemapData []opts.TreeMapNode
	for ruleName, tags := range params.Data {
		for _, val := range tags {
			intValue, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			treemapData = append(treemapData, opts.TreeMapNode{Name: ruleName, Value: intValue})
		}
	}

	treemap.AddSeries("Data", treemapData)

	return renderChartToFile(treemap, params.Output)
}

type GraphChartGenerator struct{}

func (gcg *GraphChartGenerator) Generate(params ChartParams) error {

	graph := charts.NewGraph()

	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	nodes := make([]opts.GraphNode, 0)
	links := make([]opts.GraphLink, 0)

	added := make(map[string]bool)
	for ruleName, tags := range params.Data {

		if !added[ruleName] {

			node := opts.GraphNode{Name: ruleName, Tooltip: &opts.Tooltip{Show: true}}
			nodes = append(nodes, node)
			added[ruleName] = true
		}

		for _, tag := range tags {
			if tag != ruleName {
				links = append(links, opts.GraphLink{Source: ruleName, Target: tag})
			}

			if !added[tag] {

				node := opts.GraphNode{Name: tag, Tooltip: &opts.Tooltip{Show: true}}
				nodes = append(nodes, node)
				added[tag] = true
			}
		}
	}

	graph.AddSeries("Data", nodes, links)

	return renderChartToFile(graph, params.Output)
}

type TreeChartGenerator struct{}

func (tcg *TreeChartGenerator) Generate(params ChartParams) error {
	tree := charts.NewTree()

	tree.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	tagWeights := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagWeights[tag] += 1
		}
	}

	var treeData []opts.TreeData
	added := make(map[string]bool)
	for _, tags := range params.Data {
		for _, tag := range tags {
			if !added[tag] {
				weight := tagWeights[tag]
				treeData = append(treeData, opts.TreeData{Name: tag, Value: weight})
				added[tag] = true
			}
		}
	}

	tree.AddSeries("Data", treeData)

	return renderChartToFile(tree, params.Output)
}
