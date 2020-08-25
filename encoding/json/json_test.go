package json

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Json", func() {
	type A struct {
		A int `json:"a"`
	}
	It("test MarshalString", func() {
		s := MarshalString(A{1})
		Expect(s).Should(Equal(`{"a":1}`))
	})
	It("test MarshalIndentString", func() {
		s := MarshalIndentString(A{1})
		Expect(s).Should(Equal(`{
    "a": 1
}`))
	})
})
