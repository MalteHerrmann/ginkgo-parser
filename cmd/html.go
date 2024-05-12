package cmd

import (
	"fmt"

	parserHTML "github.com/MalteHerrmann/ginkgo-parser/html"
	"github.com/MalteHerrmann/ginkgo-parser/report"
)

// ConvertGinkgoReportToHTML parses the given Ginkgo BDD JSON report and writes
// the resulting HTML document to the given export file.
func ConvertGinkgoReportToHTML(jsonFile, htmlFile string) error {
	reports, err := report.GetReportsFromJSON(jsonFile)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w\n", err)
	}

	nestedContents, err := report.BuildNestedContents(reports)
	if err != nil {
		return fmt.Errorf("error building nested contents: %w\n", err)
	}

	htmlString, err := parserHTML.RenderTemplateNew(nestedContents)
	if err != nil {
		return fmt.Errorf("error rendering HTML: %w\n", err)
	}

	return parserHTML.ExportHTML(htmlFile, htmlString)
}
