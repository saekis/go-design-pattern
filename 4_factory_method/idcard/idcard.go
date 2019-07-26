package idcard

import (
	"fmt"

	"github.com/saekis/go-design-pattern/4_factory_method/framework"
)

type IdCard struct {
	owner string
}

func NewIdCard(owner string) framework.Product {
	return &IdCard{owner: owner}
}

func (ic *IdCard) Use() string {
	return fmt.Sprintf("Use %s's card", ic.owner)
}

func (ic *IdCard) GetOwner() string {
	return ic.owner
}
