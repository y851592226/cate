package httpreq

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

type Proxy func(*http.Request) (*url.URL, error)
type Dial func(network, addr string) (net.Conn, error)
type DialTLS func(network, addr string) (net.Conn, error)
type ClientOptionFunc func(*Client) error
type CheckRedirect func(*http.Request, []*http.Request) error
type DialContext func(ctx context.Context, network, addr string) (net.Conn, error)

func SetClientTimeout(t time.Duration) ClientOptionFunc {
	return func(c *Client) error {
		c.client.Timeout = t
		return nil
	}
}

func SetClientDebug(debug, debugBody bool) ClientOptionFunc {
	return func(c *Client) error {
		c.debug = debug
		c.debugBody = debugBody
		return nil
	}
}

func AddClientRequestOptionFunc(requestOptionFunc RequestOptionFunc) ClientOptionFunc {
	return func(c *Client) error {
		c.requestOptionFuncs = append(c.requestOptionFuncs, requestOptionFunc)
		return nil
	}
}

func SetClientTransport(transport http.RoundTripper) ClientOptionFunc {
	return func(c *Client) error {
		c.client.Transport = transport
		return nil
	}
}

func SetClientCheckRedirect(checkRedirect CheckRedirect) ClientOptionFunc {
	return func(c *Client) error {
		c.client.CheckRedirect = checkRedirect
		return nil
	}
}

func SetClientCookieJar(jar http.CookieJar) ClientOptionFunc {
	return func(c *Client) error {
		c.client.Jar = jar
		return nil
	}
}

func SetClientProxyURL(URL string) ClientOptionFunc {
	return func(c *Client) error {
		proxyURL, err := url.Parse(URL)
		if err != nil {
			return err
		}
		proxy := http.ProxyURL(proxyURL)
		return SetClientProxy(proxy)(c)
	}

}

func SetClientProxy(proxy Proxy) ClientOptionFunc {
	return func(c *Client) error {
		t, ok := c.client.Transport.(*http.Transport)
		if !ok {
			return fmt.Errorf("unsupport set poroxy Transport:%T", c.client.Transport)
		}
		t.Proxy = proxy
		return nil
	}
}

func SetClientDial(dial Dial) ClientOptionFunc {
	return func(c *Client) error {
		t, ok := c.client.Transport.(*http.Transport)
		if !ok {
			return fmt.Errorf("unsupport set Dial Transport:%T", c.client.Transport)
		}
		t.Dial = dial //nolint
		return nil
	}
}

func SetClientDialContext(dialContext DialContext) ClientOptionFunc {
	return func(c *Client) error {
		t, ok := c.client.Transport.(*http.Transport)
		if !ok {
			return fmt.Errorf("unsupport set DialContext Transport:%T", c.client.Transport)
		}
		t.DialContext = dialContext
		return nil
	}
}

func SetClientDialTLS(dialTLS DialTLS) ClientOptionFunc {
	return func(c *Client) error {
		t, ok := c.client.Transport.(*http.Transport)
		if !ok {
			return fmt.Errorf("unsupport set DialTLS Transport:%T", c.client.Transport)
		}
		t.DialTLS = dialTLS
		return nil
	}
}

func AddClientMiddlewares(mws ...Middleware) ClientOptionFunc {
	return func(c *Client) error {
		c.mws = append(c.mws, mws...)
		return nil
	}
}
