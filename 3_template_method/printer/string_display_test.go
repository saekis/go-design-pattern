package printer

import "testing"

func TestStrOpen(t *testing.T) {
	cd := NewStringDisplay("string")
	expected := "+--+"
	if actual := cd.Open(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestStrPrint(t *testing.T) {
	cd := NewStringDisplay("string")
	expected := "|string|"
	if actual := cd.Print(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestStrClose(t *testing.T) {
	cd := NewStringDisplay("string")
	expected := "+--+"
	if actual := cd.Close(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
