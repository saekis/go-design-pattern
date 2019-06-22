package bookshelf

import "testing"

func Test_GetName(t *testing.T) {
	expected := "sample book name"
	book := NewBook(expected)
	if actual := book.GetName(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
