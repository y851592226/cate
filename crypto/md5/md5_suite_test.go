package md5

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMd5(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Md5 Suite")
}
