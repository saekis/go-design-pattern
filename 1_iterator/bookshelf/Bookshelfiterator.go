package bookshelf

type bookshelfIterator struct {
	bookshelf Aggregate
	index     int
}

func NewBookshelfIterator(bookshelf Aggregate) Iterator {
	return &bookshelfIterator{bookshelf: bookshelf, index: 0}
}

func (it *bookshelfIterator) HasNext() bool {
	return it.index < it.bookshelf.GetLength()
}

func (it *bookshelfIterator) Next() Product {
	book := it.bookshelf.GetByIndex(it.index)
	it.index++
	return book
}
