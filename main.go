// ginkgo-parser is a tool to parse the contents of a Ginkgo BDD test report.
// An appropriate report file can be generated using `ginkgo --json-report=FILEPATH`.
//
// It converts the specification into a markdown file, which list all run specs
// in a nested and human readable format.
//
// Optionally, an HTML file can be generated.
package main

import (
	"github.com/MalteHerrmann/ginkgo-parser/cmd"
)

func main() {
	cmd.Execute()
}
