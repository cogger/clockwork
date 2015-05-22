package automaton

import (
	"github.com/cogger/clockwork/spring"
	"golang.org/x/net/context"
	"gopkg.in/cogger/cogger.v1"
	"gopkg.in/cogger/cogger.v1/cogs"
	"gopkg.in/cogger/cogger.v1/limiter"
	"gopkg.in/cogger/cogger.v1/order"
)

//Automaton interface implements a basic cog and the ability resolve to what order to load cogs in.
type Automaton interface {
	cogger.Cog
	Wind(context.Context, spring.Spring) error
}

//New creates a new Automaton from the default automaton
func New(springs spring.Assessembly) Automaton {
	return &defaultAutomaton{
		springs: springs,
		cog:     cogs.ReturnErr(ErrUnbuildable),
		limit:   nil,
	}
}

type defaultAutomaton struct {
	springs spring.Assessembly
	cog     cogger.Cog
	limit   limiter.Limit
}

func (da *defaultAutomaton) Wind(ctx context.Context, sprng spring.Spring) error {
	sprngs, err := da.springs.Order(sprng)
	if err != nil {
		return err
	}

	da.cog = order.Series(ctx, sprngs.ToCogs()...)
	return nil
}

func (da *defaultAutomaton) Do(ctx context.Context) chan error {
	da.cog.SetLimit(da.limit)

	return da.cog.Do(ctx)
}

func (da *defaultAutomaton) SetLimit(limit limiter.Limit) cogger.Cog {
	da.limit = limit

	return da
}
