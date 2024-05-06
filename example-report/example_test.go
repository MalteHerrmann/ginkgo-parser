package example_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestExample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Example Suite")
}

var _ = Describe("Example", func() {
	It("should be true", func() {
		Expect(true).To(BeTrue())
	})

	It("should be false", func() {
		Expect(false).To(BeFalse())
	})

	DescribeTable("Nested table", func(a, b int) {
		Expect(a).To(Equal(b))
	},
		Entry("1 is 1", 1, 1),
		Entry("1 is 2", 1, 2),
	)
})
