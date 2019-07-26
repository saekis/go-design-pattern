package main

import (
	"fmt"

	"github.com/saekis/go-design-pattern/4_factory_method/framework"

	"github.com/saekis/go-design-pattern/4_factory_method/idcard"
)

func main() {
	factory := framework.Factory{}
	creater := idcard.NewIdCardFactory()
	card1 := factory.Create(creater, "たろう")
	card2 := factory.Create(creater, "はなこ")
	card3 := factory.Create(creater, "たかし")

	fmt.Println(card1.Use())
	fmt.Println(card2.Use())
	fmt.Println(card3.Use())
}
