package printer

import "testing"

func TestCharaOpen(t *testing.T) {
	cd := NewCharDisplay("character")
	expected := "<<"
	if actual := cd.Open(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCharaPrint(t *testing.T) {
	cd := NewCharDisplay("character")
	expected := "character"
	if actual := cd.Print(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestCharaClose(t *testing.T) {
	cd := NewCharDisplay("character")
	expected := ">>"
	if actual := cd.Close(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
