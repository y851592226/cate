package lrucache

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Lrucache", func() {
	It("test Lrucache", func() {
		opt := Options{
			Capacity: 3,
			OnDeleted: func(key, value interface{}) {
				println(key, value)
			},
		}
		cache, err := NewCacher(&opt)
		Expect(err).ShouldNot(HaveOccurred())
		cache.Set("1", 1)
		cache.Set("2", 2)
		cache.Set("3", 3)
		cache.Set("4", 4)
		value, ok := cache.Get("2")
		Expect(ok).Should(Equal(true))
		Expect(value).Should(Equal(2))
		cache.Set("5", 5)
		value, ok = cache.Get("3")
		Expect(ok).Should(Equal(false))
		Expect(value).Should(BeNil())
		cache.Set("6", 2)
		cache.Get("5")
		cache.Set("7", 3)
		cache.Set("8", 4)
		Expect(cache.Len()).Should(Equal(3))
	})

	It("test get int", func() {
		opt := Options{
			Capacity: 3,
		}
		cache, err := NewCacher(&opt)
		Expect(err).ShouldNot(HaveOccurred())
		cache.Set("a", 1)
		v, ok := cache.GetInt("a")
		Expect(v).Should(Equal(1))
		Expect(ok).Should(Equal(true))
		v = cache.GetIntDefault("b", 2)
		Expect(v).Should(Equal(2))
	})
	It("test get string", func() {
		opt := Options{
			Capacity: 3,
		}
		cache, err := NewCacher(&opt)
		Expect(err).ShouldNot(HaveOccurred())
		cache.Set("a", "1")
		v, ok := cache.GetString("a")
		Expect(v).Should(Equal("1"))
		Expect(ok).Should(Equal(true))
		v = cache.GetStringDefault("b", "2")
		Expect(v).Should(Equal("2"))
	})
})
