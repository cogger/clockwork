package automaton

import (
	"github.com/cogger/clockwork/spring"
	"github.com/cogger/cogger"
	"golang.org/x/net/context"
)

//Wind auto creates and returns the optizmized cog
func Wind(ctx context.Context, mainspring spring.Spring, springs spring.Assessembly) (cogger.Cog, error) {
	automaton := New(springs)
	err := automaton.Wind(ctx, mainspring)
	return automaton, err
}
