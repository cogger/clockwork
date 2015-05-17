package automaton

import (
	"github.com/cogger/clockwork/spring"
	"github.com/cogger/cogger"
	"golang.org/x/net/context"
)

func Wind(ctx context.Context, mainspring spring.Spring, springs spring.Assessembly) (cogger.Cog, error) {
	automaton := New(springs)
	err := automaton.Wind(ctx, mainspring)
	return automaton, err
}
