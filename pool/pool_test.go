package pool

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"context"
	"time"
)

var i int

func newV() (interface{}, error) {
	i++
	return i, nil
}

var _ = Describe("Pool", func() {
	opt := Options{
		NewFunc:     newV,
		MaxSize:     2,
		MaxIdleSize: 1,
		MaxIdlelime: time.Second * 1,
	}
	p := NewPool(&opt)
	It("test pool Get", func() {
		value, err := p.Get(context.Background())
		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).Should(BeEquivalentTo(1))
		p.Put(value, err)
		value, err = p.Get(context.Background())
		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).Should(BeEquivalentTo(1))
		time.Sleep(time.Second * 1)
		value, err = p.Get(context.Background())
		Expect(err).ShouldNot(HaveOccurred())
		Expect(value).Should(BeEquivalentTo(2))
		ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second/10)
		defer cancelFunc()
		value, err = p.Get(ctx)
		Expect(value).Should(BeNil())
		Expect(err).Should(HaveOccurred())
	})

})
