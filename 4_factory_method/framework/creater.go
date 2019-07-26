package framework

type Creater interface {
	CreateProduct(string) Product
	RegisterProduct(Product)
}
