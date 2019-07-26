package framework

type Factory struct {
}

func (*Factory) Create(creater Creater, owner string) Product {
	p := creater.CreateProduct(owner)
	creater.RegisterProduct(p)
	return p
}
