package base64

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBase64(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Base64 Suite")
}
