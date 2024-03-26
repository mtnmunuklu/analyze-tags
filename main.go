package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mtnmunuklu/analyze-tags/sigma"
	"github.com/mtnmunuklu/analyze-tags/yara"
)

var (
	filePath    string
	fileContent string
	showHelp    bool
	outputJSON  bool
	outputPath  string
	version     bool
	useSigma    bool
	useYara     bool
)

func init() {
	flag.StringVar(&filePath, "filepath", "", "Name or path of the file or directory to read")
	flag.StringVar(&fileContent, "filecontent", "", "Base64-encoded content of the file or directory to read")
	flag.BoolVar(&showHelp, "help", false, "Show usage")
	flag.BoolVar(&outputJSON, "json", false, "Output results in JSON format")
	flag.StringVar(&outputPath, "output", "", "Output directory for writing files")
	flag.BoolVar(&version, "version", false, "Show version information")
	flag.BoolVar(&useSigma, "sigma", false, "Use Sigma rules")
	flag.BoolVar(&useYara, "yara", false, "Use Yara rules")
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

func formatJSONResult(identifier string, tags []string) []byte {

	type JSONResult struct {
		Name string   `json:"Name"`
		Tags []string `json:"Tags"`
	}

	jsonResult := JSONResult{
		Name: identifier,
		Tags: tags,
	}

	jsonData, err := json.MarshalIndent(jsonResult, "", "  ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return nil
	}

	return jsonData
}

func main() {

	if !useSigma && !useYara {
		fmt.Println("Please provide either --sigma or --yara flag to specify the type of rules.")
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

	for _, fileContent := range fileContents {
		if useSigma {
			sigmaRule, err := sigma.ParseRule(fileContent)
			if err != nil {
				fmt.Println("Error parsing rule:", err)
				continue
			}

			var output string

			if outputJSON {
				jsonResult := formatJSONResult(sigmaRule.Title, sigmaRule.Tags)
				output = string(jsonResult)
			} else {
				output = "Name: " + sigmaRule.Title + " Tags: " + strings.Join(sigmaRule.Tags, " ")
			}

			if outputPath != "" {

				outputFilePath := filepath.Join(outputPath, fmt.Sprintf("%s.json", sigmaRule.Title))

				err := os.WriteFile(outputFilePath, []byte(output), 0644)
				if err != nil {
					fmt.Println("Error writing output to file:", err)
					continue
				}

				fmt.Printf("Output for rule '%s' written to file: %s\n", sigmaRule.Title, outputFilePath)
			} else {
				fmt.Printf("%s", output)
			}
		} else if useYara {
			yaraRuleSet, err := yara.ParseByte(fileContent)
			if err != nil {
				fmt.Println("Error parsing rule:", err)
				continue
			}

			for _, yaraRule := range yaraRuleSet.Rules {

				var output string

				if outputJSON {
					jsonResult := formatJSONResult(yaraRule.Identifier, yaraRule.Tags)
					output = string(jsonResult)
				} else {
					output = "Name: " + yaraRule.Identifier + " Tags: " + strings.Join(yaraRule.Tags, " ")
				}

				if outputPath != "" {

					outputFilePath := filepath.Join(outputPath, fmt.Sprintf("%s.json", yaraRule.Identifier))

					err := os.WriteFile(outputFilePath, []byte(output), 0644)
					if err != nil {
						fmt.Println("Error writing output to file:", err)
						continue
					}

					fmt.Printf("Output for rule '%s' written to file: %s\n", yaraRule.Identifier, outputFilePath)
				} else {
					fmt.Printf("%s", output)
				}
			}
		}
	}
}
