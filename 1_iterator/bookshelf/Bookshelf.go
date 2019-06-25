package bookshelf

type bookshelf struct {
	books []Product
	last  int
}

func NewBookShelf() *bookshelf {
	return &bookshelf{books: []Product{}, last: 0}
}

func (bs *bookshelf) GetByIndex(index int) Product {
	return bs.books[index]
}

func (bs *bookshelf) AppendBook(b Product) {
	bs.books = append(bs.books, b)
	bs.last++
}

func (bs *bookshelf) GetLength() int {
	return bs.last
}

func (bs *bookshelf) Iterator() Iterator {
	return NewBookshelfIterator(bs)
}
