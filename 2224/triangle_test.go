package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	got := area()
	expected := 6

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func TestPerimeter(t *testing.T) {
	got := perimeter()
	expected := 12

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}
