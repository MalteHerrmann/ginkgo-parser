package html_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/MalteHerrmann/ginkgo-parser/html"
	"github.com/MalteHerrmann/ginkgo-parser/report"
)

func TestHTML(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HTML Rendering Suite")
}

var _ = Describe("HTML", func() {
	It("should render the template", func() {
		r, err := report.GetReportsFromJSON("../report/testdata/example_report.json")
		Expect(err).ToNot(HaveOccurred(), "error reading JSON file")

		reportContents, err := report.BuildNestedContents(r)
		Expect(err).ToNot(HaveOccurred(), "failed to build nested contents")

		htmlString, err := html.RenderTemplateNew(reportContents)
		Expect(err).ToNot(HaveOccurred(), "error rendering template")
		Expect(htmlString).ToNot(BeEmpty(), "rendered template is empty")
	})
})
