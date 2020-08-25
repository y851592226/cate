package freejson

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFreejson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Freejson Suite")
}
