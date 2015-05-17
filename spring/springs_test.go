package spring_test

import (
	"fmt"
	"reflect"

	. "github.com/cogger/clockwork/spring"
	"github.com/cogger/cogger"
	"github.com/cogger/cogger/cogs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Springs", func() {
	coggerInterface := reflect.TypeOf((*cogger.Cog)(nil)).Elem()

	It("should return an array of cogs", func() {
		var sprngs Springs
		count := 10
		for i := 0; i < count; i++ {
			sprngs = append(sprngs, New(fmt.Sprintf("spring-$d", i), cogs.NoOp()))
		}
		cs := sprngs.ToCogs()
		Expect(cs).To(HaveLen(count))
		for _, cog := range cs {
			Expect(reflect.TypeOf(cog).Implements(coggerInterface)).To(BeTrue())
		}

	})
})
