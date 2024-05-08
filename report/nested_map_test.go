package report_test

import (
	"github.com/MalteHerrmann/ginkgo-parser/report"
	ginkgotypes "github.com/onsi/ginkgo/v2/types"

	//nolint:revive // dot imports are okay for ginkgo
	. "github.com/onsi/ginkgo/v2"
	//nolint:revive // dot imports are okay for ginkgo
	. "github.com/onsi/gomega"
)

var _ = Describe("Building the nested map", Ordered, func() {
	var (
		reports []ginkgotypes.Report
		spec    report.NestedContents
	)

	BeforeAll(func() {
		// Load the example report from the testdata directory
		var err error
		reports, err = report.GetReportsFromJSON("testdata/example_report.json")
		Expect(err).To(BeNil(), "expected no error parsing the report")
	})

	It("should build the nested map without error", func() {
		var err error
		spec, err = report.BuildNestedContents(reports)
		Expect(err).To(BeNil(), "expected no error building the spec map")
		Expect(spec).ToNot(BeNil(), "expected a non-nil spec map")
		Expect(spec.HasFailed).To(BeFalse(), "expected no failed specs")
		Expect(spec.HasSkipped).To(BeFalse(), "expected no skipped specs")
		Expect(spec.Contents).To(HaveKey("Example"), "expected a different top-level key")
	})

	It("should have the correct hierarchy", func() {
		Expect(spec.Contents["Example"].Contents).To(HaveKey("it should be true"), "expected key to be present")
		Expect(spec.Contents["Example"].Contents).To(HaveKey("it should be false"), "expected key to be present")

		Expect(spec.Contents["Example"].Contents).To(HaveKey("Nested table"), "expected key to be present")
		Expect(spec.Contents["Example"].Contents["Nested table"].Contents).To(
			HaveKey("it 1 is 1"),
			"expected key to be present",
		)

		Expect(spec.Contents["Example"].Contents).To(
			HaveKey("skipped tests"),
			"expected key to be present",
		)
		Expect(spec.Contents["Example"].Contents["skipped tests"].Contents).To(
			HaveKey("it should be skipped"),
			"expected key to be present",
		)
	})
})
