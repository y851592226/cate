package rand

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Rand", func() {
	InitSeed()
	It("test choiseN", func() {
		r := choiseN(100, 16)
		Ω(r).To(HaveLen(16))
	})
	It("test String", func() {
		r := String(NumbersLetters, 10)
		Ω(r).To(HaveLen(10))
		fmt.Println(r)
	})
	It("test BetweenFloat64", func() {
		r := BetweenFloat64(-0.1, 0.1)
		Ω(r).To(BeNumerically("<", 10))
		Ω(r).To(BeNumerically(">", -10))
		fmt.Println(r)
	})
	It("test ChoiceNFloat64", func() {
		r := ChoiceNFloat64([]float64{0, 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9}, 5)
		Ω(r).To(HaveLen(5))
		fmt.Println(r)
	})
})
