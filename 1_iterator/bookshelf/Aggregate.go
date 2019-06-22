package bookshelf

type Aggregate interface {
	Iterator() Iterator
	GetLength() int
	GetByIndex(int) Product
}
