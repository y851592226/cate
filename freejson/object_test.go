package freejson

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Object", func() {
	Context("direct convert", func() {
		It("test Float64", func() {
			o := Object{}
			o["f"] = 1.1
			f := o.Float64("f")
			Ω(f).To(BeNumerically("~", 1.1, 0.01))
		})
	})
	It("test object and array", func() {
		type TO struct {
			Int    int
			String string
			Bool   bool
		}
		type T struct {
			TO TO
			TS []TO
			T  *T
		}
		t1 := T{
			TO: TO{
				Int:    1,
				String: "aa",
				Bool:   true,
			},
			TS: []TO{{
				Int:    2,
				String: "aa",
				Bool:   true,
			}, {
				Int:    3,
				String: "aa",
				Bool:   true,
			}},
		}
		t2 := T{
			TO: TO{
				Int:    11,
				String: "aa",
				Bool:   true,
			},
			TS: []TO{{
				Int:    12,
				String: "aa",
				Bool:   true,
			}, {
				Int:    13,
				String: "aa",
				Bool:   true,
			}},
		}
		t1.T = &t2
		// By(json.MarshalIndentString(t1))
		object, err := ToObject(t1)
		Ω(err).ToNot(HaveOccurred())
		Ω(object.Object("TO").AsInt("Int")).To(Equal(1))
		Ω(object.Object("TO").String("String")).To(Equal("aa"))
		Ω(object.Array("TS").AsObjectAt(1).AsInt("Int")).To(Equal(3))
		Ω(object.Object("T").Array("TS").AsObjectAt(0).AsInt("Int")).To(Equal(12))
	})

})
