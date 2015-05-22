package spring

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"gopkg.in/cogger/cogger.v1"
	"gopkg.in/cogger/cogger.v1/cogs"
	"gopkg.in/cogger/cogger.v1/limiter"
)

var _ = Describe("Spring", func() {
	var ctx = context.Background()

	It("should also act lime a cog", func() {
		sprng := New("actLikeACog", cogs.NoOp())
		_, ok := sprng.(cogger.Cog)
		Expect(ok).To(BeTrue())

		limit, err := limiter.ByCount(5)
		Expect(err).ToNot(HaveOccurred())
		Expect(sprng.SetLimit(limit)).To(Equal(sprng))

		Expect(<-sprng.Do(ctx)).ToNot(HaveOccurred())
	})

	It("should return its name", func() {
		sprng := New("knowsItsName", cogs.NoOp())
		Expect(sprng.Name()).To(Equal("knowsItsName"))
	})

	It("should return its dependies", func() {
		deps := []string{"dep1", "dep2", "dep3"}
		sprng := New("knowsItsDependies", cogs.NoOp(), deps...)
		Expect(sprng.DependsOn()).To(HaveLen(len(deps)))
		Expect(sprng.DependsOn()).To(BeEquivalentTo(deps))
	})

})
