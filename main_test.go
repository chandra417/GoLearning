package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(2, 3)
	want := 5
	if got != want {
		t.Errorf("add(2, 3) = %d; want %d", got, want)
	}
}

func TestMultiply(t *testing.T) {
	got := multiply(2, 3)
	want := 6
	if got != want {
		t.Errorf("multiply(2, 3) = %d; want %d", got, want)
	}
}
