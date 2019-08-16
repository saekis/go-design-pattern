package main

import (
	"fmt"

	"github.com/saekis/go-design-pattern/5_singleton/singleton"
)

func main() {
	for i := 1; i <= 5; i++ {
		go func() {
			singleton.GetInstance()
		}()
	}

	fmt.Scanln()
}
