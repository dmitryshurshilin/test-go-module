package test

import "testing"

func TestConvertIntToString(t *testing.T) {
	got := ConvertIntToString(10)
	want := "10"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestConvertStringToInt(t *testing.T) {
	got := ConvertStringToInt("10")
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
