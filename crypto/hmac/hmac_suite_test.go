package hmac

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHmac(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hmac Suite")
}
