package spring

import "gopkg.in/cogger/cogger.v1"

//Springs is an array of Springs
type Springs []Spring

//ToCogs converts Springs to an array of cogs
func (sprngs Springs) ToCogs() []cogger.Cog {
	cs := make([]cogger.Cog, len(sprngs))
	for i, sprng := range sprngs {
		cs[i] = sprng
	}
	return cs
}
