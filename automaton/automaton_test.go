package automaton

import (
	"fmt"
	"reflect"

	"github.com/cogger/clockwork/spring"
	"github.com/cogger/clockwork/testdata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"gopkg.in/cogger/cogger.v1"
	"gopkg.in/cogger/cogger.v1/cogs"
)

var _ = Describe("Automaton", func() {
	coggerInterface := reflect.TypeOf((*cogger.Cog)(nil)).Elem()

	for _, test := range testdata.Tests {
		func(test testdata.TestCase) {
			Context(fmt.Sprintf("when building a %s", test.Name), func() {
				assessembly := spring.NewAssessembly()
				var first spring.Spring

				BeforeEach(func() {
					for i, sd := range test.Defs {
						sprng := assessembly.Add(spring.New(sd.Name, cogs.NoOp(), sd.Deps...))
						if i == 0 {
							first = sprng
						}
					}
				})

				It("should be able to create an automaton", func() {
					automaton := New(assessembly)
					err := automaton.Wind(context.Background(), first)
					Expect(err).ToNot(HaveOccurred())
				})

				It("should implement the cogger interface", func() {
					automaton := New(assessembly)
					err := automaton.Wind(context.Background(), first)
					Expect(err).ToNot(HaveOccurred())
					Expect(reflect.TypeOf(automaton).Implements(coggerInterface)).To(BeTrue())
					Expect(<-automaton.Do(context.Background())).ToNot(HaveOccurred())
				})

				It("should implement SetLimit function", func() {
					automaton := New(assessembly)
					err := automaton.Wind(context.Background(), first)
					Expect(err).ToNot(HaveOccurred())
					limit := &mockLimit{}

					automaton.SetLimit(limit)

					ctx := context.Background()
					Expect(<-automaton.Do(ctx)).ToNot(HaveOccurred())
					Expect(limit.NextHits).To(Equal(1))
					Expect(limit.Completed).To(BeTrue())
				})

				AfterEach(func() {
					assessembly.Clear()
				})

			})

		}(test)
	}

	It("should return an error when it can not wind a spring", func() {
		a := spring.New("a", cogs.NoOp(), "b")
		b := spring.New("b", cogs.NoOp(), "a")
		assessembly := spring.NewAssessembly()
		assessembly.Add(a)
		assessembly.Add(b)

		automaton := New(assessembly)
		err := automaton.Wind(context.Background(), a)

		Expect(err).To(HaveOccurred())
		Expect(err).To(Equal(spring.ErrCircularDependency))
	})
})

type mockLimit struct {
	Completed bool
	NextHits  int
}

func (limit *mockLimit) Next(ctx context.Context) chan struct{} {
	next := make(chan struct{})
	go func() {
		limit.NextHits++
		next <- struct{}{}
	}()
	return next
}

func (limit *mockLimit) Done(ctx context.Context) {
	limit.Completed = true
}
