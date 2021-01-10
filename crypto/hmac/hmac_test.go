package hmac

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hmac", func() {
	It("test HmacMd5", func() {
		s := Md5("key", "data")
		Ω(s).To(Equal("9d5c73ef85594d34ec4438b7c97e51d8"))
	})
})
