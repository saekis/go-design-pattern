package bookshelf

import "testing"

func Test_AppendBook(t *testing.T) {
	bs := NewBookShelf()
	bs.AppendBook(&DummyProduct2{})
	if actual := bs.last; actual != 1 {
		t.Errorf("got %v\nwant %v", actual, 1)
	}

	expected := "dummy book name"
	if actual := bs.books[0].GetName(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

// ここのテスト汚いからなんとかしたい
func Test_GetByIndex(t *testing.T) {
	expected := "dummy book name"
	bs := bookshelf{[]Product{&DummyProduct2{expected}}, 1}
	if actual := bs.GetByIndex(0).GetName(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func Test_GetLength(t *testing.T) {
	bs := bookshelf{[]Product{&book{"test"}}, 1}
	if actual := bs.GetLength(); actual != 1 {
		t.Errorf("got %v\nwant %v", actual, 1)
	}
}

func Test_Iterator(t *testing.T) {
	// TODO
}

// utilみたいな場所つくって定義した方がよさそう？
type DummyProduct2 struct {
	name string
}

func (p *DummyProduct2) GetName() string {
	return "dummy book name"
}
