package crypto

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/y851592226/cate/crypto/aes"
)

var _ = Describe("Crypto", func() {
	It("test AES Encrypt and Decrypt", func() {
		type Payload struct {
			A int
		}
		key := "123"
		encoder2 := NewEncoder(key, aes.AES{})
		payload := Payload{A: 12}
		data1, err := encoder2.Encode(payload)
		Expect(err).ShouldNot(HaveOccurred())
		payload2 := &Payload{}
		err = encoder2.Decode(data1, payload2)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(payload.A).Should(Equal(payload2.A))
	})
})
