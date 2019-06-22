package bookshelf

type Iterator interface {
	HasNext() bool
	Next() Product
}
