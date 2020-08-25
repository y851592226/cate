package freejson

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array", func() {
	It("test TypeAt", func() {
		array := Array{1, "2", 3.3, time.Now(), map[string]interface{}{"a": 1}, Array{1}}
		Expect(array.IntAt(0)).To(Equal(array[0]))
		Expect(array.StringAt(1)).To(Equal(array[1]))
		Expect(array.Float64At(2)).To(Equal(array[2]))
		Expect(array.TimeAt(3)).To(Equal(array[3]))
		Expect(array.ObjectAt(4).Int("a")).To(Equal(array[4].(map[string]interface{})["a"]))
		Expect(array.ArrayAt(5).IntAt(0)).To(Equal(array[5].(Array)[0]))
	})

	It("test AsTypeAt", func() {
		array := Array{1, "2", 3.3, time.Now()}
		Expect(array.AsStringAt(0)).To(Equal("1"))
		Expect(array.AsIntAt(1)).To(Equal(2))
		Expect(array.AsStringAt(2)).To(Equal("3.3"))
		Expect(array.AsTimeAt(3)).To(Equal(array[3]))
	})

	Context("test ToTypeAt", func() {
		It("test ToTypeAt success", func() {
			array := Array{"2009-01-01T01:01:01+08:00", "123", 1.1}
			t, err := array.ToTimeAt(0)
			Expect(err).ToNot(HaveOccurred())
			Expect(t).To(Equal(time.Date(2009, time.January, 1, 1, 1, 1, 0, time.Local)))
			i32, err := array.ToIntAt(1)
			Expect(err).ToNot(HaveOccurred())
			Expect(i32).To(Equal(123))
			f64, err := array.ToFloat64At(2)
			Expect(err).ToNot(HaveOccurred())
			Expect(f64).To(Equal(1.1))
		})

		It("test ToTypeAt error", func() {
			array := Array{"2009-01-01T01:01:01+08:00", "123", 1.1}
			_, err := array.ToFloat64At(0)
			Expect(err).To(HaveOccurred())
			_, err = array.ToIntAt(2)
			Expect(err).ToNot(HaveOccurred())
			_, err = array.ToFloat64At(3)
			Expect(err).To(HaveOccurred())
			_, err = array.ToFloat64At(-1)
			Expect(err).To(HaveOccurred())
		})

	})
})
