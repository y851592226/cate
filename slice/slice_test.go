package slice

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Slice", func() {
	var items []interface{}

	BeforeEach(func() {
		items = []interface{}{1, 2, "3", 4}
	})

	It("test Contails", func() {
		Expect(Contains(items, "3")).To(BeTrue())
		Expect(Contains(items, 3)).To(BeFalse())
	})

	It("test Index", func() {
		Expect(Index(items, "3")).To(Equal(2))
		Expect(Index(items, 3)).To(Equal(-1))
	})

	It("test Reverse", func() {
		Reverse(items)
		Expect(items[0]).To(Equal(4))
		Expect(items[1]).To(Equal("3"))
	})
})
