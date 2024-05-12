package report_test

import (
	"fmt"

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

	It("should have the correct layer descriptions", func() {
		Expect(spec.Contents["Example"].LayerDescription).To(Equal("Example"))

		Expect(spec.Contents["Example"].Contents["it should be true"].LayerDescription).To(Equal("it should be true"))
		Expect(spec.Contents["Example"].Contents["it should be false"].LayerDescription).To(Equal("it should be false"))

		fmt.Println(spec.Contents["Example"].Contents["Nested table"])
		Expect(spec.Contents["Example"].Contents["Nested table"].LayerDescription).To(Equal("Nested table"))
		Expect(spec.Contents["Example"].Contents["Nested table"].Contents["it 1 is 1"].LayerDescription).To(Equal("it 1 is 1"))
		Expect(spec.Contents["Example"].Contents["Nested table"].Contents["it 1 is 2"].LayerDescription).To(Equal("it 1 is 2"))

		Expect(spec.Contents["Example"].Contents["skipped tests"].LayerDescription).To(Equal("skipped tests"))
		Expect(spec.Contents["Example"].Contents["skipped tests"].Contents["it should be skipped"].LayerDescription).To(Equal("it should be skipped"))
	})

	It("should show the correct failed and skipped states", func() {
		Expect(spec.Contents["Example"].HasFailed).To(BeTrue())
		Expect(spec.Contents["Example"].HasSkipped).To(BeTrue())

		Expect(spec.Contents["Example"].Contents["it should be true"].HasFailed).To(BeFalse())
		Expect(spec.Contents["Example"].Contents["it should be true"].HasSkipped).To(BeFalse())

		Expect(spec.Contents["Example"].Contents["it should be false"].HasFailed).To(BeFalse())
		Expect(spec.Contents["Example"].Contents["it should be false"].HasSkipped).To(BeFalse())

		Expect(spec.Contents["Example"].Contents["Nested table"].HasFailed).To(BeTrue())
		Expect(spec.Contents["Example"].Contents["Nested table"].HasSkipped).To(BeFalse())
		Expect(spec.Contents["Example"].Contents["Nested table"].Contents["it 1 is 1"].HasFailed).To(BeFalse())
		Expect(spec.Contents["Example"].Contents["Nested table"].Contents["it 1 is 1"].HasSkipped).To(BeFalse())
		Expect(spec.Contents["Example"].Contents["Nested table"].Contents["it 1 is 2"].HasFailed).To(BeTrue())
		Expect(spec.Contents["Example"].Contents["Nested table"].Contents["it 1 is 2"].HasSkipped).To(BeFalse())

		Expect(spec.Contents["Example"].Contents["skipped tests"].HasFailed).To(BeFalse())
		Expect(spec.Contents["Example"].Contents["skipped tests"].HasSkipped).To(BeTrue())
		Expect(spec.Contents["Example"].Contents["skipped tests"].Contents["it should be skipped"].HasFailed).To(BeFalse())
		Expect(spec.Contents["Example"].Contents["skipped tests"].Contents["it should be skipped"].HasSkipped).To(BeTrue())
	})
})
