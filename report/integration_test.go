package report_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/MalteHerrmann/ginkgo-parser/report"
)

func TestReport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Report Parser Integration Tests")
}

var _ = Describe("Parsing a report JSON file", func() {
	It("should contain the available functions", func() {
		reportBytes, err := os.ReadFile("testdata/example_report.json")
		Expect(err).To(BeNil(), "expected no error reading the test data")

		// TODO: spec format here really is for the example markdown output
		spec, err := report.ParseGinkgoJSON(reportBytes)
		Expect(err).To(BeNil(), "expected no error parsing the report")
		Expect(spec).To(HaveLen(1), "expected one top-level item")
		Expect(spec).To(HaveKey("Example"))
	})
})
