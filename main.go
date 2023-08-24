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
	"fmt"
	"github.com/MalteHerrmann/ginkgo-parser/report"
	"os"
)

// defaultExportName defines the default name of the export file.
const defaultExportName = "parsed_ginkgo_suite.md"

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

	err := report.ConvertGinkgoReportToMarkdown(jsonFile, exportPath)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}
