package clockwork

import (
	"fmt"
	"sort"

	"github.com/cogger/clockwork/spring"
	"github.com/cogger/cogger"
	"github.com/cogger/cogger/cogs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

var _ = Describe("Movement", func() {
	It("should be able to add springs", func() {
		sprng := spring.New("bob", cogs.NoOp())
		Expect(Add(sprng)).To(Equal(sprng))
		addedSprng, err := Get("bob")
		Expect(err).ToNot(HaveOccurred())
		Expect(addedSprng).To(Equal(sprng))
	})

	It("should be able to add cogs with names and dependencies", func() {
		sprng := AddCog("name", cogs.NoOp())
		_, ok := sprng.(spring.Spring)
		Expect(ok).To(BeTrue())
		addedSprng, err := Get("name")
		Expect(err).ToNot(HaveOccurred())
		Expect(addedSprng).To(Equal(sprng))
	})

	It("should panic when two items of the same name are added", func() {
		Expect(func() {
			spring := spring.New("bob", cogs.NoOp())
			Add(spring)
			AddCog("bob", cogs.NoOp())
		}).To(Panic())
	})

	It("should return an error when a spring movement does not exist", func() {
		_, err := Get("somename")
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(spring.ErrDoesNotExist))
	})

	It("should panic when you are using MustGet with a spring that does not exist", func() {
		Expect(func() {
			MustGet("somename")
		}).To(Panic())
	})

	It("should return a list of names when names is called", func() {
		names := make([]string, 10)

		for i := 0; i < 10; i++ {
			names[i] = fmt.Sprintf("spring-%d", i)
			AddCog(names[i], cogs.NoOp())
		}

		returnedNames := Names()
		sort.Strings(returnedNames)
		Expect(returnedNames).To(Equal(names))
	})

	It("should empty an assessemby when clear is called", func() {
		for i := 0; i < 10; i++ {
			AddCog(fmt.Sprintf("spring-%d", i), cogs.NoOp())
		}
		Expect(Names()).To(HaveLen(10))
		Clear()
		Expect(Names()).To(HaveLen(0))
	})

	It("should return a cog when getting a movement", func() {
		AddCog("bob", cogs.NoOp())
		cog, err := Wind(context.Background(), MustGet("bob"))
		Expect(err).ToNot(HaveOccurred())
		_, ok := cog.(cogger.Cog)
		Expect(ok).To(BeTrue())
	})

	AfterEach(func() {
		Clear()
	})
})
