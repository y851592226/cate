package aes

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Aes Suite")
}
