package ttlcache

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTtlcache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ttlcache Suite")
}
