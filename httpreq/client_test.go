package httpreq

import (
	"strings"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	Context("test DefaultClient", func() {

		It("test Get request", func() {
			resp, err := Get(ts.URL+"/test/get",
				AddRequestMiddlewares(ExpectStatusCodeMW(200)),
				AddRequestQueryValue("key1", "value1"),
				AddRequestQueryValue("key1", "value11"),
				AddRequestQueryValue("key2", "value2"))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.String()).Should(Equal("this is a get request"))
		})

		It("test Post request", func() {
			resp, err := Post(ts.URL+"/test/post",
				AddRequestMiddlewares(ExpectStatusCodeMW(200)),
				SetRequestBody("post"))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.String()).Should(Equal("this is a post request"))
		})

		It("test Put request", func() {
			type Request struct {
				Put string
			}
			resp, err := Put(ts.URL+"/test/put",
				AddRequestMiddlewares(ExpectStatusCodeMW(200)),
				SetRequestBody(Request{"put"}))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.String()).Should(Equal("this is a put request"))
		})

		It("test Delete request", func() {
			var m = strings.NewReader("delete")
			resp, err := Delete(ts.URL+"/test/delete",
				AddRequestMiddlewares(ExpectStatusCodeMW(200)),
				SetRequestBody(m))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.String()).Should(Equal("this is a delete request"))
		})

	})
	Context("test ClientOptionFunc", func() {

		It("test SetClientTimeout", func() {
			client, err := NewClient(SetClientTimeout(time.Second / 10))
			Expect(err).ShouldNot(HaveOccurred())
			resp, err := client.Get(ts.URL + "/test/timeout")
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})

		It("test AddClientMiddlewares", func() {
			client, err := NewClient(AddClientMiddlewares(ExpectStatusCodeMW(400)))
			Expect(err).ShouldNot(HaveOccurred())
			resp, err := client.Get(ts.URL + "/test/get")
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})

		It("test SetClientDebug", func() {
			client, err := NewClient(SetClientDebug(true, true))
			Expect(err).ShouldNot(HaveOccurred())
			resp, err := client.Post(ts.URL+"/test/post",
				SetRequestBody("post"))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.String()).Should(Equal("this is a post request"))
		})
	})
})
