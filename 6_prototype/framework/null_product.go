package framework

type nullProduct struct{}

func NewNullProduct() Product {
	return &nullProduct{}
}

func (*nullProduct) Use(string) {
	// Do nothing
}

func (*nullProduct) CreateClone() Product {
	return &nullProduct{}
}
