package convert

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Asconvert", func() {
	It("test AsFloat64", func() {
		testCases := []struct {
			In  interface{}
			Out float64
		}{
			{"1.1", 1.1},
			{"1.1.1", 0},
		}
		for _, t := range testCases {
			out := AsFloat64(t.In)
			Expect(out).To(Equal(t.Out))
		}
	})
})
