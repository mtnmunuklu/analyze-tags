package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mtnmunuklu/analyze-tags/analytics"
	"github.com/mtnmunuklu/analyze-tags/sigma"
	"github.com/mtnmunuklu/analyze-tags/yara"
)

var (
	filePath    string
	fileContent string
	showHelp    bool
	outputPath  string
	version     bool
	useSigma    bool
	useYara     bool
	outputChart bool
	chartType   string
	outputExcel bool
)

func init() {
	flag.StringVar(&filePath, "filepath", "", "Name or path of the file or directory to read")
	flag.StringVar(&fileContent, "filecontent", "", "Base64-encoded content of the file or directory to read")
	flag.BoolVar(&showHelp, "help", false, "Show usage")
	flag.BoolVar(&version, "version", false, "Show version information")
	flag.BoolVar(&useSigma, "sigma", false, "Use Sigma rules")
	flag.BoolVar(&useYara, "yara", false, "Use Yara rules")
	flag.BoolVar(&outputChart, "chart", false, "Generate chart")
	flag.StringVar(&chartType, "chartType", "", "Specify one or more chart types to generate (comma-separated). Available chart types: bar, line, scatter, pie, boxplot, heatmap, radar, funnel, wordcloud, treemap, graph, tree")
	flag.BoolVar(&outputExcel, "excel", false, "Generate excel")
	flag.StringVar(&outputPath, "output", ".", "Output directory")

	flag.Parse()

	if version {
		fmt.Println("Analyze Tags version 1.0.0")
		os.Exit(1)
	}

	if showHelp {
		printUsage()
		os.Exit(1)
	}

	if filePath == "" && fileContent == "" {
		fmt.Println("Please provide either file paths or file contents.")
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: analyze-tags -sigma/-yara -filepath <path> [flags]")
	fmt.Println("Flags:")
	flag.PrintDefaults()
	fmt.Println("Example:")
	fmt.Println("  analyze-tags -sigma/-yara -filepath /path/to/file")
}

func generateChart(data map[string][]string, chartTypes []string) {
	for i, chartType := range chartTypes {
		foundChartType, err := analytics.FindChartType(chartType)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		output := fmt.Sprintf("%s/%s%d_chart.html", outputPath, chartType, i)
		title := fmt.Sprintf("%s chart", chartType)
		params := analytics.ChartParams{
			Type:   foundChartType,
			Data:   data,
			Title:  title,
			Output: output,
		}

		generator, err := analytics.GenerateChart(params)
		if err != nil {
			fmt.Println("Error generating chart generator: ", err)
		}

		if err := generator.Generate(params); err != nil {
			fmt.Println("Error generating chart: ", err)
		}

		if _, err := os.Stat(params.Output); os.IsNotExist(err) {
			fmt.Println("Chart file was not created:", err)
		}
	}
}

func main() {
	if !useSigma && !useYara {
		fmt.Println("Please provide either --sigma or --yara flag to specify the type of rules.")
		printUsage()
		os.Exit(1)
	}

	if !outputChart && !outputExcel {
		fmt.Println("Please provide either --chart or --excel flag to specify the output type.")
		printUsage()
		os.Exit(1)
	}

	if !outputChart && chartType == "" {
		fmt.Println("Please provide the chart type.")
		printUsage()
		os.Exit(1)
	}

	fileContents := make(map[string][]byte)

	if filePath != "" {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			fmt.Println("Error getting file/directory info:", err)
			return
		}

		if fileInfo.IsDir() {
			filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					fmt.Println("Error accessing file:", err)
					return nil
				}

				if !info.IsDir() {
					content, err := os.ReadFile(path)
					if err != nil {
						fmt.Println("Error reading file:", err)
						return nil
					}
					fileContents[path] = content
				}
				return nil
			})
		} else {
			fileContents[filePath], err = os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
		}
	} else if fileContent != "" {
		lines := strings.Split(fileContent, "\n")
		if len(lines) > 1 {

			for _, line := range lines {

				decodedContent, err := base64.StdEncoding.DecodeString(line)
				if err != nil {
					fmt.Println("Error decoding base64 content:", err)
					return
				}
				fileContents[line] = decodedContent
			}
		} else {
			decodedContent, err := base64.StdEncoding.DecodeString(fileContent)
			if err != nil {
				fmt.Println("Error decoding base64 content:", err)
				return
			}
			fileContents["filecontent"] = decodedContent
		}
	}

	data := make(map[string][]string)

	for _, fileContent := range fileContents {
		if useSigma {
			sigmaRule, err := sigma.ParseRule(fileContent)
			if err != nil {
				fmt.Println("Error parsing rule:", err)
				continue
			}

			data[sigmaRule.Title] = sigmaRule.Tags

		} else if useYara {
			yaraRuleSet, err := yara.ParseByte(fileContent)
			if err != nil {
				fmt.Println("Error parsing rule:", err)
				continue
			}

			for _, yaraRule := range yaraRuleSet.Rules {
				data[yaraRule.Identifier] = yaraRule.Tags
			}
		}
	}

	if outputChart {
		chartTypes := strings.Split(chartType, ",")
		generateChart(data, chartTypes)
	}
}
