package framework

type Product interface {
	Use(string)
	CreateClone() Product
}
