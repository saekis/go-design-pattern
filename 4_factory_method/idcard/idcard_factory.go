package idcard

import "github.com/saekis/go-design-pattern/4_factory_method/framework"

type IdCardFactory struct {
	owners []string
}

func NewIdCardFactory() framework.Creater {
	return &IdCardFactory{owners: []string{}}
}

func (*IdCardFactory) CreateProduct(owner string) framework.Product {
	return NewIdCard(owner)
}

func (factory *IdCardFactory) RegisterProduct(p framework.Product) {
	factory.owners = append(factory.owners, p.GetOwner())
}
