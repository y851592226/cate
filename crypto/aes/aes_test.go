package aes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Aes", func() {
	It("test aes Encrypt and Decrypt", func() {
		key := []byte("1234567890abaaaa")
		text := []byte("1234567890abaaaa000")
		e := AES{}
		data1, err := e.Encrypt(key, text)
		Expect(err).ShouldNot(HaveOccurred())
		data2, err := e.Decrypt(key, data1)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(text).Should(Equal(data2))
	})
})
