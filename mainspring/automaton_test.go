package mainspring

import (
	"fmt"

	"github.com/cogger/clockwork/spring"
	"github.com/cogger/cogger/cogs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

type testCase struct {
	name string
	defs []springDef
	err  error
}

type springDef struct {
	Name string
	Deps []string
}

var tests = []testCase{
	testCase{
		name: "simple dependency order",
		defs: []springDef{
			springDef{Name: "a", Deps: []string{"b", "c"}},
			springDef{Name: "b"},
			springDef{Name: "c", Deps: []string{"b"}},
		},
	},
	testCase{
		name: "deep dependency order",
		defs: []springDef{
			springDef{Name: "a", Deps: []string{"b"}},
			springDef{Name: "b", Deps: []string{"c"}},
			springDef{Name: "c", Deps: []string{"d"}},
			springDef{Name: "d", Deps: []string{"e"}},
			springDef{Name: "e"},
		},
	},
	testCase{
		name: "multibranch",
		defs: []springDef{
			springDef{Name: "a", Deps: []string{"b", "c", "i"}},
			springDef{Name: "b", Deps: []string{"c", "e"}},
			springDef{Name: "c", Deps: []string{"d"}},
			springDef{Name: "d", Deps: []string{"e", "j"}},
			springDef{Name: "e", Deps: []string{"f"}},
			springDef{Name: "f", Deps: []string{"g"}},
			springDef{Name: "g"},
			springDef{Name: "h", Deps: []string{"e"}},
			springDef{Name: "i", Deps: []string{"j"}},
			springDef{Name: "j", Deps: []string{"k"}},
			springDef{Name: "k"},
		},
	},
	testCase{
		name: "circular dependency",
		defs: []springDef{
			springDef{Name: "a", Deps: []string{"b"}},
			springDef{Name: "b", Deps: []string{"c"}},
			springDef{Name: "c", Deps: []string{"d"}},
			springDef{Name: "d", Deps: []string{"e"}},
			springDef{Name: "e", Deps: []string{"c"}},
		},
		err: ErrCircularDependency,
	},
	testCase{
		name: "dependency does not exist",
		defs: []springDef{
			springDef{Name: "a", Deps: []string{"b"}},
			springDef{Name: "b", Deps: []string{"c"}},
			springDef{Name: "c", Deps: []string{"d"}},
			springDef{Name: "d", Deps: []string{"e"}},
			springDef{Name: "e", Deps: []string{"f"}},
		},
		err: spring.ErrDoesNotExist,
	},
}

var _ = Describe("Automaton", func() {
	for _, test := range tests {
		func(test testCase) {
			It(fmt.Sprintf("should be able to do a %s", test.name), func() {
				assessembly := spring.NewAssessembly()
				var first spring.Spring
				for i, sd := range test.defs {

					sprng := assessembly.Add(spring.New(sd.Name, cogs.NoOp(), sd.Deps...))
					if i == 0 {
						first = sprng
					}
				}
				_, err := Wind(context.Background(), first, assessembly)

				if test.err == nil {
					Expect(err).ToNot(HaveOccurred())
				} else {
					Expect(err).To(HaveOccurred())
					Expect(err).To(Equal(test.err))
				}

			})
		}(test)
	}

})
