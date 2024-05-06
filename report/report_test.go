package report_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/MalteHerrmann/ginkgo-parser/report"
)

func TestReport(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Report Parser Integration Tests")
}

var _ = Describe("Parsing a Ginkgo JSON file", func() {
	It("should return the expected reports", func() {
		// TODO: report format here really is for the example markdown output
		report, err := report.ParseGinkgoJSON("testdata/example_report.json")
		Expect(err).To(BeNil(), "expected no error parsing the report")
		Expect(report).To(HaveLen(1), "expected one top-level item")
		Expect(report).To(HaveKey("Example"))
	})
})
