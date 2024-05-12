package html

import (
	"errors"
	"html/template"
	"os"
	"strings"

	"github.com/MalteHerrmann/ginkgo-parser/report"
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
	t, err := template.ParseFiles("html/index.html")
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	if err := t.ExecuteTemplate(&builder, "index", reportData); err != nil {
		return "", err
	}

	return builder.String(), nil
}

// RenderTemplateNew renders the report data into an HTML document.
func RenderTemplateNew(nestedContents report.NestedContents) (string, error) {
	t, err := template.ParseFiles("html/index_new.html")
	if err != nil {
		return "", err
	}

	var builder strings.Builder
	if err = t.ExecuteTemplate(&builder, "index", nestedContents); err != nil {
		return "", err
	}

	return builder.String(), nil
}

// ExportHTML exports the given HTML string to a file.
func ExportHTML(filename, htmlString string) error {
	if !strings.HasSuffix(filename, ".html") {
		return errors.New("target must be a .html file")
	}

	return os.WriteFile(filename, []byte(htmlString), 0o600)
}
