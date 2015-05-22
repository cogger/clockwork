package automaton

import (
	"fmt"

	"github.com/cogger/clockwork/spring"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"gopkg.in/cogger/cogger.v1"
	"gopkg.in/cogger/cogger.v1/cogs"
)

var _ = Describe("Wind", func() {
	var ctx = context.Background()
	It("should return a cog when winding a spring", func() {
		start := spring.New("start", cogs.NoOp())
		assessembly := spring.NewAssessembly()
		assessembly.Add(start)

		for i := 0; i > 10; i++ {
			assessembly.Add(spring.New(fmt.Sprintf("spring-%d", i), cogs.NoOp()))
		}
		returnedCog, err := Wind(ctx, start, assessembly)
		Expect(err).ToNot(HaveOccurred())
		_, ok := returnedCog.(cogger.Cog)
		Expect(ok).To(BeTrue())
	})

	It("should return an error when it can not wind a spring", func() {
		a := spring.New("a", cogs.NoOp(), "b")
		b := spring.New("b", cogs.NoOp(), "a")
		assessembly := spring.NewAssessembly()
		assessembly.Add(a)
		assessembly.Add(b)
		_, err := Wind(ctx, a, assessembly)
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(spring.ErrCircularDependency))
	})
})
