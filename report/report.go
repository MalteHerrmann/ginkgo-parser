package report

import (
	"encoding/json"
	"fmt"
	"os"

	ginkgotypes "github.com/onsi/ginkgo/v2/types"
)

// spacesPerIndentation defines the whitespace to be used per indentation level.
const spacesPerIndentation = "  "

// ConvertGinkgoReportToMarkdown parses the given Ginkgo BDD JSON report and writes
// the resulting markdown to the given export file.
func ConvertGinkgoReportToMarkdown(jsonFile, markdownFile string) error {
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		return fmt.Errorf("error reading file: %w\n", err)
	}

	var reports []ginkgotypes.Report
	err = json.Unmarshal(file, &reports)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w\n", err)
	}

	spec := make(map[string]interface{})

	for _, report := range reports {
		for _, specReport := range report.SpecReports {
			containerHierarchy := specReport.ContainerHierarchyTexts

			var leafNodeTypeStr string
			leafNodeType := specReport.LeafNodeType
			switch leafNodeType {
			case ginkgotypes.NodeTypeIt:
				leafNodeTypeStr = "it"
			default:
				panic(fmt.Sprintf("Unknown leaf node type: %d", leafNodeType))
			}

			leafNodeText := specReport.LeafNodeText
			cleanLeafNode := leafNodeTypeStr + " " + leafNodeText
			containerHierarchy = append(containerHierarchy, cleanLeafNode)

			currentSpec := spec
			for _, item := range containerHierarchy {
				// If the item is not yet in the map, create a new map for the underlying levels
				// and add the item to the map.
				if currentSpec[item] == nil {
					currentSpec[item] = make(map[string]interface{})
				}
				currentSpec = currentSpec[item].(map[string]interface{})
			}
		}
	}

	markdownContents := buildMarkdown(spec, 0)
	err = os.WriteFile(markdownFile, []byte(markdownContents), 0o644)
	if err != nil {
		return fmt.Errorf("Error writing file: %w\n", err)
	}

	fmt.Printf("Markdown file '%s' generated successfully.\n", markdownFile)
	return nil
}

// buildMarkdown is a recursive function to build a markdown string from the given map of nested maps.
func buildMarkdown(contents map[string]interface{}, indentationLevel int) string {
	markdown := ""
	for key, value := range contents {
		if value != nil {
			markdown += fmt.Sprintf("%s- %s\n", spaces(indentationLevel), key)
			markdown += buildMarkdown(value.(map[string]interface{}), indentationLevel+1)
		} else {
			markdown += fmt.Sprintf("%s- %s\n", spaces(indentationLevel), key)
		}
	}
	return markdown
}

// spaces returns a string of n spaces, which each consist of
// the specified number of spacesPerIndentation.
func spaces(n int) string {
	spaceStr := ""
	for i := 0; i < n; i++ {
		spaceStr += spacesPerIndentation
	}
	return spaceStr
}
