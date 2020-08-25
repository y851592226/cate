package lrucache

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLrucache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lrucache Suite")
}
