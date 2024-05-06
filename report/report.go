// Package report provides functionality to parse
// a Ginkgo BDD JSON report, and convert its contents
// to a nested markdown file.
package report

import (
	"encoding/json"
	"fmt"
	"os"

	ginkgotypes "github.com/onsi/ginkgo/v2/types"
)

// GetReportsFromJSON parses the given JSON file into Ginkgo reports.
func GetReportsFromJSON(jsonFile string) ([]ginkgotypes.Report, error) {
	reportBytes, err := os.ReadFile(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var reports []ginkgotypes.Report
	if err := json.Unmarshal(reportBytes, &reports); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	return reports, nil
}
