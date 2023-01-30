package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	got := area()
	expected := 81

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func TestPerimeter(t *testing.T) {
	got := perimeter()
	expected := 36

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
