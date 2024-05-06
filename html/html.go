package html

import (
	"errors"
	"html/template"
	"os"
	"strings"

	ginkgotypes "github.com/onsi/ginkgo/v2/types"
)

// ReportData collects the data needed to render the report.
type ReportData struct {
	Reports []ginkgotypes.Report
}

// NewReportData creates a new ReportData instance from the given reports.
func NewReportData(reports []ginkgotypes.Report) ReportData {
	return ReportData{
		Reports: reports,
	}
}

// RenderTemplate renders the report data into an HTML string.
func RenderTemplate(reportData ReportData) (string, error) {
	t, err := template.New("index").ParseFiles("index.html")
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	if err := t.Execute(&builder, reportData); err != nil {
		return "", err
	}

	return builder.String(), nil
}

// ExportHTML exports the given HTML string to a file.
func ExportHTML(htmlString, filename string) error {
	if !strings.HasSuffix(filename, ".html") {
		return errors.New("target must be a .html file")
	}

	return os.WriteFile(filename, []byte(htmlString), 0o600)
}
