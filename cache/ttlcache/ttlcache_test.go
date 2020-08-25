package ttlcache

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ttlcache", func() {
	It("Test method Get", func() {
		var count int32
		var num int32
		fetcher := func(ctx context.Context, key interface{}) (interface{}, error) {
			time.Sleep(time.Millisecond * 100)
			atomic.AddInt32(&num, 1)
			return num, nil
		}
		cacher := NewCache(Options{time.Millisecond * 200, fetcher})
		wg := sync.WaitGroup{}
		wg.Add(10000)
		now := time.Now()
		for i := 0; i < 10000; i++ {
			go func() {
				for range time.Tick(time.Millisecond) {
					atomic.AddInt32(&count, 1)
					n, _ := cacher.Get(context.Background(), "hello")
					if n.(int32) > 5 {
						break
					}
				}
				wg.Done()
			}()
		}
		wg.Wait()
		By(time.Since(now).String())

		Expect(count > 100000).Should(Equal(true))
		Expect(num).Should(Equal(int32(6)))
	})
})
