package automaton

import (
	"github.com/cogger/clockwork/spring"
	"golang.org/x/net/context"
	"gopkg.in/cogger/cogger.v1"
)

//Wind auto creates and returns the optizmized cog
func Wind(ctx context.Context, mainspring spring.Spring, springs spring.Assessembly) (cogger.Cog, error) {
	automaton := New(springs)
	err := automaton.Wind(ctx, mainspring)
	return automaton, err
}
