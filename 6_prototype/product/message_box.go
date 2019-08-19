package product

import (
	"fmt"

	"github.com/saekis/go-design-pattern/6_prototype/framework"
)

type messageBox struct {
	charDecorated string
}

func NewMessageBox(str string) framework.Product {
	return &messageBox{str}
}

func (mb *messageBox) Use(s string) {
	len := len(s)
	for i := 0; i < len+4; i++ {
		fmt.Print(mb.charDecorated)
	}
	fmt.Println("")
	fmt.Println(fmt.Sprintf("%s %s %s", mb.charDecorated, s, mb.charDecorated))
	for i := 0; i < len+4; i++ {
		fmt.Print(mb.charDecorated)
	}
	fmt.Println("")
}

func (mb *messageBox) CreateClone() framework.Product {
	return &messageBox{mb.charDecorated}
}
