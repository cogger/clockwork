package mainspring

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMainspring(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mainspring Suite")
}
