package spring

import (
	"github.com/cogger/cogger"
	"github.com/cogger/cogger/limiter"
	"golang.org/x/net/context"
)

//Spring implements a basic cog interface, the ability to get a name and
type Spring interface {
	cogger.Cog
	Name() string
	DependsOn() []string
}

//New creates a spring from a name, cog and dependency names and default implementation
func New(name string, cog cogger.Cog, dependsOn ...string) Spring {
	return &defaultSpring{
		name:    name,
		cog:     cog,
		depends: dependsOn,
	}
}

type defaultSpring struct {
	name    string
	cog     cogger.Cog
	depends []string
}

func (spring *defaultSpring) Name() string {
	return spring.name
}

func (spring *defaultSpring) DependsOn() []string {
	return spring.depends
}

func (spring *defaultSpring) Do(ctx context.Context) chan error {
	return spring.cog.Do(ctx)
}

func (spring *defaultSpring) SetLimit(limit limiter.Limit) cogger.Cog {
	spring.cog.SetLimit(limit)
	return spring
}
