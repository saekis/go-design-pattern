package product

import (
	"fmt"

	"github.com/saekis/go-design-pattern/6_prototype/framework"
)

type underlinePen struct {
	charUl string
}

func NewUnderlinePen(str string) framework.Product {
	return &underlinePen{str}
}

func (up *underlinePen) Use(str string) {
	len := len(str)
	fmt.Println("\"" + str + "\"")
	for i := 0; i < len; i++ {
		fmt.Print(up.charUl)
	}
	fmt.Println("")
}

func (up *underlinePen) CreateClone() framework.Product {
	return &underlinePen{up.charUl}
}
