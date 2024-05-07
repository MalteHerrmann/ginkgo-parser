package report

import (
	"fmt"

	parserHTML "github.com/MalteHerrmann/ginkgo-parser/html"
)

// ConvertGinkgoReportToHTML parses the given Ginkgo BDD JSON report and writes
// the resulting HTML document to the given export file.
func ConvertGinkgoReportToHTML(jsonFile, htmlFile string) error {
	reports, err := GetReportsFromJSON(jsonFile)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %w\n", err)
	}

	htmlString, err := parserHTML.RenderTemplate(parserHTML.NewReportData(reports))
	if err != nil {
		return fmt.Errorf("error rendering HTML: %w\n", err)
	}

	return parserHTML.ExportHTML(htmlFile, htmlString)
}
