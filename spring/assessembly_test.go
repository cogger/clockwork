package spring

import (
	"fmt"
	"sort"

	"github.com/cogger/cogger/cogs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Assessembly", func() {
	var assessembly = NewAssessembly()

	It("should be able to add springs", func() {
		sprng := New("bob", cogs.NoOp())
		Expect(assessembly.Add(sprng)).To(Equal(sprng))
		addedSprng, err := assessembly.Get("bob")
		Expect(err).ToNot(HaveOccurred())
		Expect(addedSprng).To(Equal(sprng))
	})

	It("should panic when two items of the same name are added", func() {
		Expect(func() {
			assessembly.Add(New("bob", cogs.NoOp()))
			assessembly.Add(New("bob", cogs.NoOp()))
		}).To(Panic())
	})

	It("should return an error when a spring movement does not exist", func() {
		_, err := assessembly.Get("somename")
		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(ErrDoesNotExist))
	})

	It("should return a list of names when names is called", func() {
		names := make([]string, 10)

		for i := 0; i < 10; i++ {
			names[i] = fmt.Sprintf("spring-%d", i)
			assessembly.Add(New(names[i], cogs.NoOp()))
		}

		returnedNames := assessembly.Names()
		sort.Strings(returnedNames)
		Expect(returnedNames).To(Equal(names))
	})

	It("should empty an assessemby when clear is called", func() {
		for i := 0; i < 10; i++ {
			assessembly.Add(New(fmt.Sprintf("spring-%d", i), cogs.NoOp()))
		}
		Expect(assessembly.Names()).To(HaveLen(10))
		assessembly.Clear()
		Expect(assessembly.Names()).To(HaveLen(0))
	})

	AfterEach(func() {
		assessembly.Clear()
	})

})
