package ratelimit

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ratelimit", func() {
	It("test BlockingGet1", func() {
		tb := NewTokenBucket(1, 100, time.Second/100)
		begin := time.Now()
		for i := 1; i < 50; i++ {
			tb.BlockingGet()
		}
		Expect(time.Since(begin) < time.Millisecond*1500).Should(Equal(true))
	})
})
