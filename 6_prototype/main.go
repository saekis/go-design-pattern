package main

import (
	"github.com/saekis/go-design-pattern/6_prototype/framework"
	"github.com/saekis/go-design-pattern/6_prototype/product"
)

func main() {
	manager := framework.NewManager()
	up := product.NewUnderlinePen("~")
	mb1 := product.NewMessageBox("*")
	mb2 := product.NewMessageBox("/")
	manager.Register("Strong message", up)
	manager.Register("Warning box", mb1)
	manager.Register("Slash box", mb2)

	p1 := manager.Create("Strong message")
	p1.Use("Hello World.")
	p2 := manager.Create("Warning box")
	p2.Use("Hello World.")
	p3 := manager.Create("Slash box")
	p3.Use("Hello World.")
}
