package httpreq

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Request", func() {
	Context("test RequestOptionFunc", func() {

		It("test SetRequestRetryTimes", func() {
			var count int
			Counter := func() Middleware {
				return func(next EndPoint) EndPoint {
					return func(req *Request) (*Response, error) {
						count++
						return next(req)
					}
				}
			}
			resp, err := Get("https://www.noexists.com",
				SetRequestRetryTimes(3),
				SetRequestDebug(true, false),
				AddRequestMiddlewares(Counter()))
			Expect(count).Should(Equal(4))
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})

		It("test SetRequestFormValues", func() {
			resp, err := Post(ts.URL+"/test/postform",
				AddRequestFormValue("key1", "value1"),
				AddRequestFormValue("key1", "value11"),
				AddRequestFormValue("key2", "value2"),
			)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.String()).Should(Equal("key1=value1&key1=value11&key2=value2"))
		})

		It("test AddRequestCookie", func() {
			resp, err := Get(ts.URL+"/test/requestcookie",
				AddRequestCookie("user", "cate"))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.String()).Should(Equal("user=cate"))
		})
	})
})
