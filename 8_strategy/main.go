package main

import (
	"os"

	"github.com/saekis/go-design-pattern/8_strategy/strategy"
)

func main() {
	var s strategy.Strategy

	switch os.Args[1] {
	case "alpha":
		s = &strategy.AlphaStrategy{}
	case "beta":
		s = &strategy.BetaStrategy{}
	default:
		return
	}

	Context{s}.Exec()
}
