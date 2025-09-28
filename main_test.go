package main

import (
	"fmt"
	"testing"
)

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

func TestSumOfNumbers(t *testing.T) {

	nums := []int{1, 2, 3}
	sum := 0

	for _, n := range nums {
		sum += n
	}
	if sum != 6 {
		t.Errorf("sum = %d; want 6", sum)
	}
}

func TestOddEvenNumbers(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for _, n := range nums {
		fmt.Println(oddEven(n))
	}
}
