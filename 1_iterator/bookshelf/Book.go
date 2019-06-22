package bookshelf

type book struct {
	name string
}

func NewBook(name string) *book {
	return &book{name}
}

func (b *book) GetName() string {
	return b.name
}
