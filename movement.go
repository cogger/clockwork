package clockwork

import (
	"github.com/cogger/clockwork/mainspring"
	"github.com/cogger/clockwork/spring"
	"github.com/cogger/cogger"
	"golang.org/x/net/context"
)

var assessembly = spring.NewAssessembly()

func AddCog(name string, cog cogger.Cog, dependsOn ...string) spring.Spring {
	return Add(spring.New(name, cog, dependsOn...))
}

func Add(sprng spring.Spring) spring.Spring {
	return assessembly.Add(sprng)
}

func Get(name string) (spring.Spring, error) {
	return assessembly.Get(name)
}

func MustGet(name string) spring.Spring {
	sprng, err := assessembly.Get(name)
	if err != nil {
		panic(err)
	}
	return sprng
}

func Names() []string {
	return assessembly.Names()
}

func Clear() {
	assessembly.Clear()
}

func Movement(ctx context.Context, sprng spring.Spring) (cogger.Cog, error) {
	return mainspring.Wind(ctx, sprng, assessembly)
}
