package spring

import "github.com/cogger/cogger"

type Springs []Spring

func (sprngs Springs) ToCogs() []cogger.Cog {
	cs := make([]cogger.Cog, len(sprngs))
	for i, sprng := range sprngs {
		cs[i] = sprng
	}
	return cs
}
