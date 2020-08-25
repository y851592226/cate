package convert

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Asconvertdefault", func() {
	Context("test AsFloat64Default", func() {
		It("test AsFloat64Default", func() {
			testCases := []struct {
				In      interface{}
				Default float64
				Out     float64
			}{
				{"1.1", -1, 1.1},
				{"1.1.1", -1, -1},
			}
			for _, t := range testCases {
				out := AsFloat64Default(t.In, t.Default)
				Expect(out).To(Equal(t.Out))
			}
		})

	})
})
