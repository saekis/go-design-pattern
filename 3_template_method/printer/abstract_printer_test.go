package printer

import "testing"

func TestDisplay(t *testing.T) {
	ad := AbstractPrinter{}
	expected := "open.print.close."
	if actual := ad.Display(&dummyPrinter{}); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

type dummyPrinter struct {
}

func (*dummyPrinter) Open() string {
	return "open."
}

func (*dummyPrinter) Print() string {
	return "print."
}

func (*dummyPrinter) Close() string {
	return "close."
}
