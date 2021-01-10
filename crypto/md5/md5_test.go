package md5

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Md5", func() {
	It("test MD5", func() {
		s := MD5("md5")
		Ω(s).To(Equal("1bc29b36f623ba82aaf6724fd3b16718"))
	})
})
