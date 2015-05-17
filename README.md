# clockwork

| Type               | Badge
|---------------------------------------
| **Documentation:** | [![GoDoc](https://godoc.org/github.com/cogger/clockwork?status.png)](http://godoc.org/github.com/cogger/clockwork)  
| **Build Status:**  | [![Build Status](https://travis-ci.org/cogger/clockwork.svg?branch=master)](https://travis-ci.org/cogger/clockwork)  
| **Test Coverage:** | [![Coverage Status](https://coveralls.io/repos/cogger/clockwork/badge.svg?branch=master)](https://coveralls.io/r/cogger/clockwork?branch=master)
| **License:**       | [![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)

clockwork extends base cogger to add named cogs with dependency lists.  clockwork also contains functionality to do a graph topological sort to determine what order to call the cogs in.  It currently calls all cogs in a series but is being upgraded to call all cogs that can be in parallel.

## Usage

~~~ go
package main

import (
	"github.com/cogger/cogger/cogs"
	"github.com/cogger/clockwork"
	"github.com/cogger/clockwork/spring"
	"golang.org/x/net/context"
)

func init() {
	keySprng := spring.New("a",cogs.NoOp,"b","c")
	
	clockwork.Add(keySprng)
	clockwork.AddCog("b",cogs.NoOp())
	clockwork.AddCog("c",cogs.NoOp(),"b")

	cog, err := clockwork.Wind(context.Background(),keySprng)
	...	
}
~~~