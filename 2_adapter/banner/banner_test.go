package banner

import "testing"

func TestWithParen(t *testing.T) {
	expected := "(string)"
	banner := Banner{"string"}
	if actual := banner.WithParen(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestWithAster(t *testing.T) {
	expected := "*string*"
	banner := Banner{"string"}
	if actual := banner.WithAster(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
