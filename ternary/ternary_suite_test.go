package ternary

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTernary(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ternary Suite")
}
