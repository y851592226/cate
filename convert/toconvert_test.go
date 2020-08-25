package convert

import (
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ui8  uint8   = 255   //nolint
	ui16 uint16  = 255   //nolint
	ui32 uint32  = 255   //nolint
	ui64 uint64  = 255   //nolint
	i8   int8    = -127  //nolint
	i16  int16   = -127  //nolint
	i32  int32   = -127  //nolint
	i64  int64   = -127  //nolint
	f32  float32 = 1.25  //nolint
	f64  float64 = 1.25  //nolint
	i    int     = 12345 //nolint
	ui   uint    = 12345 //nolint
)

var _ = Describe("Toconvert", func() {
	Context("test ToTime", func() {
		It("test ToTime success", func() {
			testCases := []struct {
				In  interface{}
				Out time.Time
			}{
				{time.Date(2009, time.January, 1, 1, 1, 1, 1, time.Local), time.Date(2009, time.January, 1, 1, 1, 1, 1, time.Local)},
				{"2009-01-01 01:01:01", time.Date(2009, time.January, 1, 1, 1, 1, 0, time.Local)},
				{"2009-01-01T01:01:01+08:00", time.Date(2009, time.January, 1, 1, 1, 1, 0, time.Local)},
				{int(1230742861), time.Date(2009, time.January, 1, 1, 1, 1, 0, time.Local)},
				{int64(1230742861000000001), time.Date(2009, time.January, 1, 1, 1, 1, 1, time.Local)},
				{"2019-09-01T12:12:12Z", time.Date(2019, time.September, 1, 12, 12, 12, 0, time.UTC)},
				{"2019-09-01 12:12:12", time.Date(2019, time.September, 1, 12, 12, 12, 0, time.Local)},
				{"2019-09-01T12:12:12+08:00", time.Date(2019, time.September, 1, 12, 12, 12, 0, time.Local)},
			}
			for _, t := range testCases {
				out, err := ToTime(t.In)
				Expect(err).ToNot(HaveOccurred())
				Expect(out).To(Equal(t.Out))
			}
		})
	})
	Context("test ToString", func() {
		It("test ToString success", func() {
			testCases := []struct {
				In  interface{}
				Out string
			}{
				{true, "true"},
				{false, "false"},
				{1, "1"},
				{1.1, "1.1"},
				{time.Date(2009, time.January, 1, 1, 1, 1, 1, time.Local), "2009-01-01T01:01:01+08:00"},
				{errors.New("this is a error"), "this is a error"},
			}
			for _, t := range testCases {
				out, err := ToString(t.In)
				Expect(err).ToNot(HaveOccurred())
				Expect(out).To(Equal(t.Out))
			}

		})

	})
	Context("Test ToInt64", func() {
		It("test ToInt64 success", func() {
			testCases := []struct {
				In  interface{}
				Out int64
			}{
				{"1", 1},
				{"0", 0},
				{true, 1},
				{false, 0},
				{f64, 1},
				{f32, 1},
				{"-12345", -12345},
			}
			for _, t := range testCases {
				out, err := ToInt64(t.In)
				Expect(err).ToNot(HaveOccurred())
				Expect(out).To(Equal(t.Out))
			}
		})

	})

	Context("Test ToFloat64", func() {
		It("test ToFloat64 success", func() {
			testCases := []struct {
				In  interface{}
				Out float64
			}{
				{"1", 1},
				{"0", 0},
				{true, 1},
				{false, 0},
				{f64, 1.25},
				{f32, 1.25},
				{"-12345", -12345},
			}
			for _, t := range testCases {
				out, err := ToFloat64(t.In)
				Expect(err).ToNot(HaveOccurred())
				Expect(out).To(Equal(t.Out))
			}

		})
	})

})
