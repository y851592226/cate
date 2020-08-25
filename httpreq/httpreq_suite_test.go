package httpreq

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	ts *httptest.Server
)

type FooStruct struct {
	Foo string `json:"foo" xml:"foo"`
}

type EFooStruct struct {
	Foo int `json:"foo" xml:"foo"`
}

func TestHttpreq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Httpreq Suite")
}

var _ = BeforeSuite(func() {
	setupRoute()
})

func setupRoute() {
	gin.SetMode(gin.ReleaseMode)
	handler := gin.New()
	handler.GET("/test/timeout", func(c *gin.Context) {
		time.Sleep(time.Second)
		c.String(200, "hello")
	})
	handler.GET("/test/get", func(c *gin.Context) {
		c.String(200, "this is a get request")
	})
	handler.POST("/test/post", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		Expect(err).ShouldNot(HaveOccurred())
		if string(data) != "post" {
			c.String(400, "bad request")
		} else {
			c.String(200, "this is a post request")
		}
	})
	handler.PUT("/test/put", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		Expect(err).ShouldNot(HaveOccurred())
		if string(data) != `{"Put":"put"}` {
			c.String(400, "bad request")
		} else {
			c.String(200, "this is a put request")
		}
	})
	handler.DELETE("/test/delete", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		Expect(err).ShouldNot(HaveOccurred())
		if string(data) != "delete" {
			c.String(400, "bad request")
		} else {
			c.String(200, "this is a delete request")
		}
	})
	handler.POST("/test/postform", func(c *gin.Context) {
		if c.ContentType() == "application/x-www-form-urlencoded" {
			data, err := ioutil.ReadAll(c.Request.Body)
			Expect(err).ShouldNot(HaveOccurred())
			c.String(200, string(data))
		} else {
			c.String(400, "bad request")
		}

	})
	handler.GET("/test/requestcookie", func(c *gin.Context) {
		user, err := c.Cookie("user")
		if err != nil {
			c.String(403, "forbidden")
		}
		c.String(200, "user="+user)
	})
	handler.Any("/test/response/error", func(c *gin.Context) {
		c.Header("Content-Length", "100")
		c.String(200, "hello world")
	})
	handler.Any("/test/response/success", func(c *gin.Context) {
		c.Header("server", "gin")
		c.SetCookie("package", "github.com/y851592226/cate/httpreq", 1000, "", "", false, true)
		c.String(200, "welcome")
	})
	handler.Any("/test/response/bind/json", func(c *gin.Context) {
		c.Header("server", "gin")
		c.SetCookie("package", "github.com/y851592226/cate/httpreq", 1000, "", "", false, true)
		c.JSON(200, FooStruct{"bar-json"})
	})
	handler.Any("/test/response/bind/xml", func(c *gin.Context) {
		c.Header("server", "gin")
		c.SetCookie("package", "github.com/y851592226/cate/httpreq", 1000, "", "", false, true)
		c.XML(200, FooStruct{"bar-xml"})
	})
	ts = httptest.NewServer(handler)
}
