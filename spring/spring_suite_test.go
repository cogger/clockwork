package spring

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSpring(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Spring Suite")
}
