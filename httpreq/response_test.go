package httpreq

import (
	"net/url"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Response", func() {

	Context("test readBody Response", func() {
		var resp *Response

		BeforeEach(func() {
			var err error
			resp, err = Get(ts.URL + "/test/response/success")
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should get body correctly", func() {
			Expect(resp.Body()).Should(Equal([]byte(`welcome`)))
		})

		It("should status equal 200 ok", func() {
			Expect(resp.Status()).Should(Equal("200 OK"))
		})

		It("should StatusCode equal 200", func() {
			Expect(resp.StatusCode()).Should(Equal(200))
		})

		It("should get string correctly", func() {
			Expect(resp.String()).Should(Equal(`welcome`))
		})

		It("should get response header correctly", func() {
			header := resp.ResponseHeader()
			Expect(header).ShouldNot(BeNil())
		})

		It("should get header correctly", func() {
			server := resp.Header("server")
			Expect(server).Should(Equal("gin"))
		})

		It("should get cookies correctly", func() {
			cookies := resp.Cookies()
			Expect(len(cookies)).Should(Equal(1))
			cookie := cookies[0]
			Expect(cookie.Name).Should(Equal("package"))
			Expect(url.QueryUnescape(cookie.Value)).Should(Equal("github.com/y851592226/cate/httpreq"))
		})
	})

	Context("test bind Response", func() {
		var respJSON *Response
		var respXML *Response

		BeforeEach(func() {
			var err error
			respJSON, err = Get(ts.URL + "/test/response/bind/json")
			Expect(err).ShouldNot(HaveOccurred())
			respXML, err = Get(ts.URL + "/test/response/bind/xml")
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("should bind json success", func() {
			fooStruct := FooStruct{}
			err := respJSON.BindJSON(&fooStruct)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(fooStruct.Foo).Should(Equal("bar-json"))
		})

		It("should bind xml success", func() {
			fooStruct := FooStruct{}
			err := respXML.BindXML(&fooStruct)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(fooStruct.Foo).Should(Equal("bar-xml"))
		})

		It("should bind success", func() {
			fooStruct := FooStruct{}
			err := respJSON.Bind(&fooStruct)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(fooStruct.Foo).Should(Equal("bar-json"))

			err = respXML.Bind(&fooStruct)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(fooStruct.Foo).Should(Equal("bar-xml"))
		})

		It("should bind failed", func() {
			fooStruct := EFooStruct{}
			err := respJSON.Bind(&fooStruct)
			Expect(err).Should(HaveOccurred())
		})
	})

})
