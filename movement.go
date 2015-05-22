package clockwork

import (
	"github.com/cogger/clockwork/automaton"
	"github.com/cogger/clockwork/spring"
	"golang.org/x/net/context"
	"gopkg.in/cogger/cogger.v1"
)

var assessembly = spring.NewAssessembly()

//AddCog addes a cog to the clockwork assessembly
func AddCog(name string, cog cogger.Cog, dependsOn ...string) spring.Spring {
	return Add(spring.New(name, cog, dependsOn...))
}

//Add addes a spring to the clockwork assessembly
func Add(sprng spring.Spring) spring.Spring {
	return assessembly.Add(sprng)
}

//Get gets a spring from the clockwork assessembly by name
func Get(name string) (spring.Spring, error) {
	return assessembly.Get(name)
}

//MustGet panics if it can not get the spring from the clockwork assessembly
func MustGet(name string) spring.Spring {
	sprng, err := assessembly.Get(name)
	if err != nil {
		panic(err)
	}
	return sprng
}

//Names returns the list of names of the springs in the clockwork assessembly
func Names() []string {
	return assessembly.Names()
}

//Clear clears the clockwork assessembly of all springs
func Clear() {
	assessembly.Clear()
}

//Wind resolves all the dependencies and returns a single cog.
func Wind(ctx context.Context, sprng spring.Spring) (cogger.Cog, error) {
	return automaton.Wind(ctx, sprng, assessembly)
}
