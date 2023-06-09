// ginkgo-parser is a tool to parse the contents of a Ginkgo BDD test report.
// An appropriate report file can be generated using `ginkgo --json-report=FILEPATH`.
//
// It converts the specification into a markdown file, which holds a run specs in a nested and
// human readable format.
//
// Usage:
//
//	go run github.com/MalteHerrmann/ginkgo-parser GINKGO_REPORT [EXPORT_PATH]
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	// defaultExportName defines the default name of the export file.
	defaultExportName = "parsed_ginkgo_suite.md"
	// spacesPerIndentation defines the whitespace to be used per indentation level.
	spacesPerIndentation = "  "
)

// buildMarkdown recursively builds a markdown string from the given map.
func buildMarkdown(d map[string]interface{}, indentationLevel int) string {
	markdown := ""
	for key, value := range d {
		if value != nil {
			markdown += fmt.Sprintf("%s- %s\n", spaces(indentationLevel), key)
			markdown += buildMarkdown(value.(map[string]interface{}), indentationLevel+1)
		} else {
			markdown += fmt.Sprintf("%s- %s\n", spaces(indentationLevel), key)
		}
	}
	return markdown
}

// convertGinkgoReportToMarkdown parses the given Ginkgo BDD JSON report and writes
// the resulting markdown to the given export file.
func convertGinkgoReportToMarkdown(jsonFile, markdownFile string) {
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return
	}

	data := []map[string]interface{}{}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Printf("Error parsing JSON: %s\n", err)
		return
	}

	spec := make(map[string]interface{})

	for _, testsuite := range data {
		for _, specReport := range testsuite["SpecReports"].([]interface{}) {
			// TODO: implement types instead of using interface{}
			containerHierarchy := specReport.(map[string]interface{})["ContainerHierarchyTexts"].([]interface{})
			leafNodeType := specReport.(map[string]interface{})["LeafNodeType"].(string)
			switch leafNodeType {
			case "It":
				leafNodeType = "it"
			default:
				panic("Unknown leaf node type: " + leafNodeType)
			}

			leafNodeText := specReport.(map[string]interface{})["LeafNodeText"].(string)
			cleanLeafNode := leafNodeType + " " + leafNodeText
			containerHierarchy = append(containerHierarchy, cleanLeafNode)

			currentSpec := spec
			for _, item := range containerHierarchy {
				if currentSpec[item.(string)] == nil {
					currentSpec[item.(string)] = make(map[string]interface{})
				}
				currentSpec = currentSpec[item.(string)].(map[string]interface{})
			}
		}
	}

	markdownContents := buildMarkdown(spec, 0)
	err = os.WriteFile(markdownFile, []byte(markdownContents), 0o644)
	if err != nil {
		fmt.Printf("Error writing file: %s\n", err)
		return
	}

	fmt.Printf("Markdown file '%s' generated successfully.\n", markdownFile)
}

// spaces returns a string of n spaces.
func spaces(n int) string {
	spaceStr := ""
	for i := 0; i < n; i++ {
		spaceStr += spacesPerIndentation
	}
	return spaceStr
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: ginkgo-parser GINKGO_REPORT [EXPORT_PATH]")
		return
	}

	jsonFile := os.Args[1]
	if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
		fmt.Printf("File '%s' not found.\n", jsonFile)
		return
	}

	exportPath := defaultExportName
	if len(os.Args) == 3 {
		exportPath = os.Args[2]
	}

	convertGinkgoReportToMarkdown(jsonFile, exportPath)
}
