package analytics

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// ChartType represents the type of chart to be used
type ChartType string

const (
	// BarChart represents the bar chart type
	BarChart ChartType = "bar"
	// LineChart represents the line chart type
	LineChart ChartType = "line"
	// ScatterPlot represents the scatter plot type
	ScatterPlot ChartType = "scatter"
	// PieChart represents the pie chart type
	PieChart ChartType = "pie"
	// BoxPlotChart represents the box plot chart type
	BoxPlotChart ChartType = "boxplot"
	// HeatmapChart represents the heatmap chart type
	HeatmapChart ChartType = "heatmap"
	// RadarChart represents the radar chart type
	RadarChart ChartType = "radar"
	// GaugeChart represents the gauge chart type
	GaugeChart ChartType = "gauge"
	// FunnelChart represents the funnel chart type
	FunnelChart ChartType = "funnel"
	// SankeyChart represents the sankey chart type
	SankeyChart ChartType = "sankey"
	// WordCloudChart represents the word cloud chart type
	WordCloudChart ChartType = "wordcloud"
	// TreemapChart represents the treemap chart type
	TreemapChart ChartType = "treemap"
	// GraphChart represents the graph chart type
	GraphChart ChartType = "graph"
	// KlineChart represents the kline chart type
	KlineChart ChartType = "kline"
	// ParallelChart represents the parallel chart type
	ParallelChart ChartType = "parallel"
	// SunburstChart represents the sunburst chart type
	SunburstChart ChartType = "sunburst"
	// SurfaceChart represents the surface chart type
	SurfaceChart ChartType = "surface"
	// ThemeRiverChart represents the theme river chart type
	ThemeRiverChart ChartType = "themeriver"
	// TreeChart represents the tree chart type
	TreeChart ChartType = "tree"
)

// ChartParams contains the parameters required for the chart
type ChartParams struct {
	Type   ChartType           // The type of chart
	Data   map[string][]string // Rule names and tags
	Title  string              // The title of the chart
	Output string              // The output path for the chart file
}

// renderChartToFile function renders the specified chart to a file.
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

// GenerateChart function returns the appropriate chart generator based on the specified chart type.
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
	case GaugeChart:
		return &GaugeChartGenerator{}, nil
	case FunnelChart:
		return &FunnelChartGenerator{}, nil
	case SankeyChart:
		return &SankeyChartGenerator{}, nil
	case WordCloudChart:
		return &WordCloudChartGenerator{}, nil
	case TreemapChart:
		return &TreemapChartGenerator{}, nil
	case GraphChart:
		return &GraphChartGenerator{}, nil
	case KlineChart:
		return &KlineChartGenerator{}, nil
	case ParallelChart:
		return &ParallelChartGenerator{}, nil
	case SunburstChart:
		return &SunburstChartGenerator{}, nil
	case SurfaceChart:
		return &SurfaceChartGenerator{}, nil
	case ThemeRiverChart:
		return &ThemeRiverChartGenerator{}, nil
	case TreeChart:
		return &TreeChartGenerator{}, nil
	default:
		return nil, fmt.Errorf("unsupported chart type: %s", params.Type)
	}
}

// ChartGenerator represents an interface for chart generation functionality
type ChartGenerator interface {
	Generate(params ChartParams) error // Generate chart function
}

// BarChartGenerator is an implementation for generating bar charts
type BarChartGenerator struct{}

// Generate implements the function to generate bar charts
func (bcg *BarChartGenerator) Generate(params ChartParams) error {
	// Create a bar chart
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Calculate tag distribution
	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	// Create data for the chart
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

	// Write the chart to a file
	return renderChartToFile(bar, params.Output)
}

// LineChartGenerator is an implementation for generating line charts
type LineChartGenerator struct{}

// Generate implements the function to generate line charts
func (lcg *LineChartGenerator) Generate(params ChartParams) error {
	// Create a line chart
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Calculate tag distribution
	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	// Create data for the chart
	var xAxisData []string
	var lineData []opts.LineData
	for tag, count := range tagCounts {
		xAxisData = append(xAxisData, tag)
		lineData = append(lineData, opts.LineData{Value: count})
	}

	line.SetXAxis(xAxisData).
		AddSeries("Count", lineData)

	// Write the chart to a file
	return renderChartToFile(line, params.Output)
}

// ScatterPlotGenerator is an implementation for generating scatter plots
type ScatterPlotGenerator struct{}

// Generate implements the function to generate scatter plots
func (spg *ScatterPlotGenerator) Generate(params ChartParams) error {
	// Create a scatter plot
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Calculate tag distribution
	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	// Create data for the chart
	var scatterData []opts.ScatterData
	var xAxisData []string
	for tag, count := range tagCounts {
		xAxisData = append(xAxisData, tag)
		scatterData = append(scatterData, opts.ScatterData{Value: []interface{}{tag, count}})
	}

	// Add series to scatter plot
	scatter.SetXAxis(xAxisData).AddSeries("Count", scatterData)

	// Write the chart to a file
	return renderChartToFile(scatter, params.Output)
}

// PieChartGenerator is an implementation for generating pie charts
type PieChartGenerator struct{}

// Generate implements the function to generate pie charts
func (pcg *PieChartGenerator) Generate(params ChartParams) error {
	// Create a pie chart
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Calculate tag distribution and total count
	tagCounts := make(map[string]int)
	totalCount := 0
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
			totalCount++
		}
	}

	// Create data for the chart
	var pieData []opts.PieData
	for tag, count := range tagCounts {
		// Calculate percentage
		percentage := float64(count) / float64(totalCount) * 100
		label := fmt.Sprintf("{b}: %d (%.2f%%)", count, percentage)
		pieData = append(pieData, opts.PieData{Name: tag, Value: count, Label: &opts.Label{Show: true, Formatter: label}})
	}

	pie.AddSeries("Count", pieData)

	// Write the chart to a file
	return renderChartToFile(pie, params.Output)
}

// BoxPlotGenerator is an implementation for generating box plots
type BoxPlotGenerator struct{}

// Generate implements the function to generate box plots
func (bpg *BoxPlotGenerator) Generate(params ChartParams) error {
	// Create a box plot
	boxplot := charts.NewBoxPlot()
	boxplot.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Calculate tag distribution
	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	// Create data for the chart
	var xAxisData []string
	var boxplotData []opts.BoxPlotData
	for tag, count := range tagCounts {
		xAxisData = append(xAxisData, tag)
		boxplotData = append(boxplotData, opts.BoxPlotData{Name: tag, Value: count})
	}

	boxplot.SetXAxis(xAxisData)
	boxplot.AddSeries("Count", boxplotData)

	// Write the chart to a file
	return renderChartToFile(boxplot, params.Output)
}

// HeatmapGenerator is an implementation for generating heatmaps
type HeatmapGenerator struct{}

// Generate function creates heatmaps based on provided parameters
func (hmg *HeatmapGenerator) Generate(params ChartParams) error {
	// Create a new heatmap instance
	heatmap := charts.NewHeatMap()

	// Set global options for the heatmap
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

	// Calculate tag counts
	tagCounts := make(map[string]int)
	for _, tags := range params.Data {
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	// Generate heatmap data
	var data []opts.HeatMapData
	for rule, tags := range params.Data {
		for _, tag := range tags {
			// Set the numerical value as the count of the respective tag
			count := tagCounts[tag]
			data = append(data, opts.HeatMapData{Value: [3]interface{}{tag, rule, count}})
		}
	}

	// Add series to the heatmap
	heatmap.AddSeries("Count", data)

	// Write the chart to a file
	return renderChartToFile(heatmap, params.Output)
}

// RadarChartGenerator is an implementation for generating radar charts
type RadarChartGenerator struct{}

// Generate implements the function to generate radar charts
func (rcg *RadarChartGenerator) Generate(params ChartParams) error {
	// Create a radar chart
	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare radar indicator data
	indicators := make([]*opts.Indicator, 0)
	for rule := range params.Data {
		indicators = append(indicators, &opts.Indicator{Name: rule})
	}

	// Prepare radar series data
	seriesData := make([]opts.RadarData, 0)
	for _, tags := range params.Data {
		var rowData []interface{}
		for _, tag := range tags {
			rowData = append(rowData, tag)
		}
		seriesData = append(seriesData, opts.RadarData{Value: rowData})
	}

	// Set radar options
	radar.SetGlobalOptions(charts.WithRadarComponentOpts(opts.RadarComponent{
		Indicator: indicators,
	}))

	// Add series to the radar
	radar.AddSeries("Data", seriesData)

	// Write the chart to a file
	return renderChartToFile(radar, params.Output)
}

// GaugeChartGenerator is an implementation for generating gauge charts
type GaugeChartGenerator struct{}

// Generate implements the function to generate gauge charts
func (gcg *GaugeChartGenerator) Generate(params ChartParams) error {
	// Create a gauge chart
	gauge := charts.NewGauge()
	gauge.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the gauge chart
	var data []opts.GaugeData
	for tag, tags := range params.Data {
		var value float32 = 0
		for _, val := range tags {
			// Convert string values to float32 for data processing
			floatValue, err := strconv.ParseFloat(val, 32)
			if err != nil {
				return err
			}
			value += float32(floatValue)
		}
		data = append(data, opts.GaugeData{Name: tag, Value: value})
	}

	// Add series data to the gauge chart
	gauge.AddSeries("Data", data)

	// Write the chart to a file
	return renderChartToFile(gauge, params.Output)
}

// FunnelChartGenerator is an implementation for generating funnel charts
type FunnelChartGenerator struct{}

// Generate implements the function to generate funnel charts
func (fcg *FunnelChartGenerator) Generate(params ChartParams) error {
	// Create a funnel chart
	funnel := charts.NewFunnel()
	funnel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the funnel chart
	var data []opts.FunnelData
	for tag, tags := range params.Data {
		var value float32 = 0
		for _, val := range tags {
			// Convert string values to float32 for data processing
			floatValue, err := strconv.ParseFloat(val, 32)
			if err != nil {
				return err
			}
			value += float32(floatValue)
		}
		data = append(data, opts.FunnelData{Name: tag, Value: value})
	}

	// Add series data to the funnel chart
	funnel.AddSeries("Data", data)

	// Write the chart to a file
	return renderChartToFile(funnel, params.Output)
}

// SankeyChartGenerator is an implementation for generating sankey charts
type SankeyChartGenerator struct{}

// Generate implements the function to generate sankey charts
func (scg *SankeyChartGenerator) Generate(params ChartParams) error {
	// Create a sankey chart
	sankey := charts.NewSankey()
	sankey.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the sankey chart
	var links []opts.SankeyLink
	var nodes []opts.SankeyNode
	nodeID := make(map[string]int)
	count := 0
	for _, tags := range params.Data {
		for i := 0; i < len(tags)-1; i++ {
			sourceID, ok := nodeID[tags[i]]
			if !ok {
				sourceID = count
				count++
				nodeID[tags[i]] = sourceID
				nodes = append(nodes, opts.SankeyNode{Name: tags[i]})
			}
			targetID, ok := nodeID[tags[i+1]]
			if !ok {
				targetID = count
				count++
				nodeID[tags[i+1]] = targetID
				nodes = append(nodes, opts.SankeyNode{Name: tags[i+1]})
			}
			links = append(links, opts.SankeyLink{Source: fmt.Sprintf("%d", sourceID), Target: fmt.Sprintf("%d", targetID)})
		}
	}

	// Set data for the sankey chart
	sankey.AddSeries("sankey", nodes, links).
		SetSeriesOptions(
			charts.WithLineStyleOpts(opts.LineStyle{
				Color:     "source",
				Curveness: 0.5,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)

	// Write the chart to a file
	return renderChartToFile(sankey, params.Output)
}

// WordCloudChartGenerator is an implementation for generating word cloud charts
type WordCloudChartGenerator struct{}

// Generate implements the function to generate word cloud charts
func (wcg *WordCloudChartGenerator) Generate(params ChartParams) error {
	// Create a word cloud chart
	wordCloud := charts.NewWordCloud()
	wordCloud.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the word cloud chart
	var data []opts.WordCloudData
	for ruleName, tags := range params.Data {
		for _, val := range tags {
			// Convert string values to float32 for data processing
			floatValue, err := strconv.ParseFloat(val, 32)
			if err != nil {
				return err
			}
			data = append(data, opts.WordCloudData{Name: ruleName, Value: floatValue})
		}
	}

	// Add series data to the word cloud chart
	wordCloud.AddSeries("Data", data)

	// Write the chart to a file
	return renderChartToFile(wordCloud, params.Output)
}

// TreemapChartGenerator is an implementation for generating treemap charts
type TreemapChartGenerator struct{}

// Generate implements the function to generate treemap charts
func (tcg *TreemapChartGenerator) Generate(params ChartParams) error {
	// Create a treemap chart
	treemap := charts.NewTreeMap()
	treemap.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the treemap chart
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

	// Add series data to the treemap chart
	treemap.AddSeries("Data", treemapData)

	// Write the chart to a file
	return renderChartToFile(treemap, params.Output)
}

// GraphChartGenerator is an implementation for generating graph charts
type GraphChartGenerator struct{}

// Generate implements the function to generate graph charts
func (gcg *GraphChartGenerator) Generate(params ChartParams) error {
	// Create a new graph chart
	graph := charts.NewGraph()
	// Set global options for the graph chart
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the graph chart
	nodes := make([]opts.GraphNode, 0)
	links := make([]opts.GraphLink, 0)
	for ruleName, tags := range params.Data {
		for _, val := range tags {
			floatValue, err := strconv.ParseFloat(val, 32)
			if err != nil {
				return err
			}
			// Create a new graph node and append it to the nodes slice
			nodes = append(nodes, opts.GraphNode{Name: ruleName, Value: float32(floatValue)})
			// Create links between nodes based on the graph type
			// Here, we are assuming a simple undirected graph
			for _, tag := range tags {
				if tag != ruleName {
					links = append(links, opts.GraphLink{Source: ruleName, Target: tag})
				}
			}
		}
	}

	// Add series data to the graph chart
	graph.AddSeries("Data", nodes, links)

	// Render the chart to a file
	return renderChartToFile(graph, params.Output)
}

// KlineChartGenerator is an implementation for generating kline charts
type KlineChartGenerator struct{}

// Generate implements the function to generate kline charts
func (kcg *KlineChartGenerator) Generate(params ChartParams) error {
	// Create a new kline chart
	kline := charts.NewKLine()
	// Set global options for the kline chart
	kline.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the kline chart
	klineData := make([]opts.KlineData, 0)
	for date, values := range params.Data {
		open, _ := strconv.ParseFloat(values[0], 64)
		close, _ := strconv.ParseFloat(values[1], 64)
		low, _ := strconv.ParseFloat(values[2], 64)
		high, _ := strconv.ParseFloat(values[3], 64)
		klineData = append(klineData, opts.KlineData{Value: [4]float64{open, close, low, high}, Name: date})
	}

	// Add series data to the kline chart
	kline.AddSeries("Kline", klineData)

	// Render the chart to a file
	return renderChartToFile(kline, params.Output)
}

// ParallelChartGenerator is an implementation for generating parallel charts
type ParallelChartGenerator struct{}

// Generate implements the function to generate parallel charts
func (pcg *ParallelChartGenerator) Generate(params ChartParams) error {
	// Create a new parallel chart
	parallel := charts.NewParallel()
	// Set global options for the parallel chart
	parallel.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the parallel chart
	parallelData := make([]opts.ParallelData, 0)
	for _, tags := range params.Data {
		// Convert string values to float64 for data processing
		values := make([]interface{}, len(tags))
		for i, val := range tags {
			floatValue, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			values[i] = floatValue
		}
		parallelData = append(parallelData, opts.ParallelData{Value: values})
	}

	// Add series data to the parallel chart
	parallel.AddSeries("Data", parallelData)

	// Render the chart to a file
	return renderChartToFile(parallel, params.Output)
}

// SunburstChartGenerator is an implementation for generating sunburst charts
type SunburstChartGenerator struct{}

// Generate implements the function to generate sunburst charts
func (scg *SunburstChartGenerator) Generate(params ChartParams) error {
	// Create a new sunburst chart
	sunburst := charts.NewSunburst()
	// Set global options for the sunburst chart
	sunburst.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the sunburst chart
	sunburstData := make([]opts.SunBurstData, 0)
	for ruleName, tags := range params.Data {
		for _, val := range tags {
			// Convert string values to float64 for data processing
			floatValue, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			sunburstData = append(sunburstData, opts.SunBurstData{Name: ruleName, Value: floatValue})
		}
	}

	// Add series data to the sunburst chart
	sunburst.AddSeries("Data", sunburstData)

	// Render the chart to a file
	return renderChartToFile(sunburst, params.Output)
}

// SurfaceChartGenerator is an implementation for generating surface charts
type SurfaceChartGenerator struct{}

// Generate implements the function to generate surface charts
func (scg *SurfaceChartGenerator) Generate(params ChartParams) error {
	// Create a new surface chart
	surface := charts.NewSurface3D()
	// Set global options for the surface chart
	surface.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the surface chart
	var surfaceData []opts.Chart3DData
	for ruleName, tags := range params.Data {
		for _, val := range tags {
			// Convert string values to float64 for data processing
			floatValue, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			surfaceData = append(surfaceData, opts.Chart3DData{Value: []interface{}{ruleName, val, floatValue}})
		}
	}

	// Add series data to the surface chart
	surface.AddSeries("Data", surfaceData)

	// Render the chart to a file
	return renderChartToFile(surface, params.Output)
}

// ThemeRiverChartGenerator is an implementation for generating theme river charts
type ThemeRiverChartGenerator struct{}

// Generate implements the function to generate theme river charts
func (trcg *ThemeRiverChartGenerator) Generate(params ChartParams) error {
	// Create a new theme river chart
	themeRiver := charts.NewThemeRiver()
	// Set global options for the theme river chart
	themeRiver.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the theme river chart
	var themeRiverData []opts.ThemeRiverData
	for ruleName, tags := range params.Data {
		for _, val := range tags {
			// Convert string values to float64 for data processing
			floatValue, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			themeRiverData = append(themeRiverData, opts.ThemeRiverData{Name: ruleName, Value: floatValue})
		}
	}

	// Add series data to the theme river chart
	themeRiver.AddSeries("Data", themeRiverData)

	// Render the chart to a file
	return renderChartToFile(themeRiver, params.Output)
}

// TreeChartGenerator is an implementation for generating tree charts
type TreeChartGenerator struct{}

// Generate implements the function to generate tree charts
func (tcg *TreeChartGenerator) Generate(params ChartParams) error {
	// Create a new tree chart
	tree := charts.NewTree()
	// Set global options for the tree chart
	tree.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: params.Title,
		}),
	)

	// Prepare data for the tree chart
	var treeData []opts.TreeData
	for tag, tags := range params.Data {
		for _, val := range tags {
			// Convert string values to float64 for data processing
			floatValue, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			// Convert float64 to int
			intValue := int(floatValue)
			treeData = append(treeData, opts.TreeData{Name: tag, Value: intValue})
		}
	}

	// Add series data to the tree chart
	tree.AddSeries("Data", treeData)

	// Render the chart to a file
	return renderChartToFile(tree, params.Output)
}
