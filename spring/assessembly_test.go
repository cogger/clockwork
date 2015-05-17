package spring

import (
	"fmt"
	"sort"

	"github.com/cogger/clockwork/testdata"
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

	Context("when ordering springs", func() {
		for _, test := range testdata.Tests {
			func(test testdata.TestCase) {
				It(fmt.Sprintf("should be able to do a %s", test.Name), func() {

					var first Spring
					for i, sd := range test.Defs {

						sprng := assessembly.Add(New(sd.Name, cogs.NoOp(), sd.Deps...))
						if i == 0 {
							first = sprng
						}
					}

					path, err := assessembly.Order(first)

					Expect(err).ToNot(HaveOccurred())
					names := []string{}
					for _, sprng := range path {
						names = append(names, sprng.Name())
					}
					Expect(names).To(Equal(test.Expected))
				})
			}(test)
		}

		It("should fail on circular dependencies", func() {
			sprng := assessembly.Add(New("a", cogs.NoOp(), "b"))
			assessembly.Add(New("b", cogs.NoOp(), "c"))
			assessembly.Add(New("c", cogs.NoOp(), "d"))
			assessembly.Add(New("d", cogs.NoOp(), "e"))
			assessembly.Add(New("e", cogs.NoOp(), "c"))

			_, err := assessembly.Order(sprng)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrCircularDependency))
		})

		It("should fail when a dependencies does not exist", func() {
			sprng := assessembly.Add(New("a", cogs.NoOp(), "b"))
			assessembly.Add(New("b", cogs.NoOp(), "c"))
			assessembly.Add(New("c", cogs.NoOp(), "d"))
			assessembly.Add(New("d", cogs.NoOp(), "e"))
			assessembly.Add(New("e", cogs.NoOp(), "f"))

			_, err := assessembly.Order(sprng)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrDoesNotExist))
		})

		It("should fail when a the spring you are trying to resolve is not in the assessembly", func() {
			sprng := New("a", cogs.NoOp(), "b")
			assessembly.Add(New("b", cogs.NoOp(), "c"))
			assessembly.Add(New("c", cogs.NoOp(), "d"))
			assessembly.Add(New("d", cogs.NoOp(), "e"))
			assessembly.Add(New("e", cogs.NoOp()))

			_, err := assessembly.Order(sprng)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ErrDoesNotExist))
		})
	})

	AfterEach(func() {
		assessembly.Clear()
	})

})
