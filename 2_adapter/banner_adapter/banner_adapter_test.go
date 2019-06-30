package banner_adapter

import "testing"

func TestGetWeak(t *testing.T) {
	expected := "(string)"
	ba := NewBannerAdapter("string")
	if actual := ba.GetWeak(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}

func TestGetStrong(t *testing.T) {
	expected := "*string*"
	ba := NewBannerAdapter("string")
	if actual := ba.GetStrong(); actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
