package automaton

import "errors"

//ErrUnbuildable is returned when an order was not built
var ErrUnbuildable = errors.New("unable to build automaton")
