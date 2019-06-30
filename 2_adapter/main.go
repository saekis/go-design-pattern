package main

import (
	"fmt"

	"github.com/saekis/go-design-pattern/2_adapter/banner_adapter"
)

func main() {
	b := banner_adapter.NewBannerAdapter("Hello")
	fmt.Println(b.GetWeak())
	fmt.Println(b.GetStrong())
}
