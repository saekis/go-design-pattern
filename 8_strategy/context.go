package main

import "github.com/saekis/go-design-pattern/8_strategy/strategy"

type Context struct {
	strategy strategy.Strategy
}

func (c *Context) Exec() {
	c.strategy.Call()
}
