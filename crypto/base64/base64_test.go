package base64

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Base64", func() {
	It("test EncodeToString", func() {
		s, err := EncodeToString(1)
		Ω(err).ToNot(HaveOccurred())
		Ω(s).To(Equal("MQ=="))
		By(s)
	})
})
