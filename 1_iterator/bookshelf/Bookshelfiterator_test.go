package bookshelf

import "testing"

func Test_HasNext(t *testing.T) {
	it := NewBookshelfIterator(&DummyBookShelf{})
	if it.HasNext() != true {
		t.Error("got false\nwant true")
	}
}

func Test_Next(t *testing.T) {
	it := NewBookshelfIterator(&DummyBookShelf{})
	expected := "dummy book name"
	if actual := it.Next().GetName(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

type DummyBookShelf struct {
}

func (bs *DummyBookShelf) Iterator() Iterator {
	return &DummyIterator{}
}

func (bs *DummyBookShelf) GetLength() int {
	return 1
}

func (bs *DummyBookShelf) GetByIndex(int) Product {
	return &DummyProduct{}
}

type DummyIterator struct {
}

func (it *DummyIterator) HasNext() bool {
	return true
}

func (it *DummyIterator) Next() Product {
	return &DummyProduct{}
}

type DummyProduct struct {
}

func (p *DummyProduct) GetName() string {
	return "dummy book name"
}
