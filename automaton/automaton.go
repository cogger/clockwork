package automaton

import (
	"github.com/cogger/clockwork/spring"
	"github.com/cogger/cogger"
	"github.com/cogger/cogger/cogs"
	"github.com/cogger/cogger/limiter"
	"github.com/cogger/cogger/order"
	"golang.org/x/net/context"
)

type Automaton interface {
	cogger.Cog
	Wind(context.Context, spring.Spring) error
}

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
