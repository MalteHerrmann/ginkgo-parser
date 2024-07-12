package cmd

import (
	"fmt"
	"os"

	"github.com/MalteHerrmann/ginkgo-parser/report"
)

// spacesPerIndentation defines the whitespace to be used per indentation level.
const spacesPerIndentation = "  "

// ConvertGinkgoReportToMarkdown parses the given Ginkgo BDD JSON report and writes
// the resulting markdown to the given export file.
func ConvertGinkgoReportToMarkdown(jsonFile, markdownFile string) error {
	reports, err := report.GetReportsFromJSON(jsonFile)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w", err)
	}

	spec, err := report.BuildNestedContents(reports)
	if err != nil {
		return fmt.Errorf("error building spec map: %w", err)
	}

	markdownContents := buildMarkdown(spec, 0)
	err = os.WriteFile(markdownFile, []byte(markdownContents), 0o600)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	fmt.Printf("Markdown file '%s' generated successfully.\n", markdownFile)
	return nil
}

// buildMarkdown is a recursive function to build a markdown string from the given map of nested maps.
func buildMarkdown(spec report.NestedContents, indentationLevel int) string {
	markdown := ""
	for key, value := range spec.Contents {
		if len(value.Contents) != 0 {
			markdown += fmt.Sprintf("%s- %s\n", spaces(indentationLevel), key)
			markdown += buildMarkdown(value, indentationLevel+1)
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
