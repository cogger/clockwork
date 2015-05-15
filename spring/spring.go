package spring

import (
	"github.com/cogger/cogger"
	"github.com/cogger/cogger/limiter"
	"golang.org/x/net/context"
)

type Spring interface {
	cogger.Cog
	Name() string
	DependsOn() []string
}

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
