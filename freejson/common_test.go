package freejson

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/y851592226/cate/encoding/json"
)

var _ = Describe("Common", func() {
	Context("ToArray function", func() {
		It("test string to Array", func() {
			s := "[1,2,3]"
			array, err := ToArray(s)
			Ω(err).ToNot(HaveOccurred())
			By(json.MarshalIndentString(array))
		})
		It("test []byte to Array", func() {
			s := []byte("[1,2,3]")
			array, err := ToArray(s)
			Ω(err).ToNot(HaveOccurred())
			By(json.MarshalIndentString(array))
		})

		It("test array to Array", func() {
			s := [3]float64{1.1, 2.2, 3.3}
			array, err := ToArray(s)
			Ω(err).ToNot(HaveOccurred())
			By(json.MarshalIndentString(array))
		})

		It("test slice to Array", func() {
			s := []string{"a", "b", "c"}
			array, err := ToArray(s)
			Ω(err).ToNot(HaveOccurred())
			By(json.MarshalIndentString(array))
		})
	})

	Context("ToObject function", func() {
		It("test string to Object", func() {
			s := `{"a":"a","b":"b"}`
			o, err := ToObject(s)
			Ω(err).ToNot(HaveOccurred())
			By(json.MarshalIndentString(o))
		})
		It("test []byte to Object", func() {
			s := []byte(`{"a":"a","b":"b"}`)
			o, err := ToObject(s)
			Ω(err).ToNot(HaveOccurred())
			By(json.MarshalIndentString(o))
		})

		It("test struct to Object", func() {
			type TS struct {
				A int
				B string
			}
			s := TS{1, "a"}
			o, err := ToObject(s)
			Ω(err).ToNot(HaveOccurred())
			By(json.MarshalIndentString(o))
		})

		It("test map to Object", func() {
			s := map[string]int{
				"a": 1,
				"b": 2,
			}
			o, err := ToObject(s)
			Ω(err).ToNot(HaveOccurred())
			By(json.MarshalIndentString(o))
		})
	})

})
