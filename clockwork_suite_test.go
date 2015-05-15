package clockwork

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestClockwork(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Clockwork Suite")
}
