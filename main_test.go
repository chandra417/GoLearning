package main

import (
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

func TestOddNumbers(t *testing.T) {

	var nums = [4]int{1, 3, 5, 7}
	for _, n := range nums {
		odd := oddNumber(n)
		if !odd {
			t.Errorf("odd = %d; want %t", n, true)
		}
	}
}

func TestWordCount(t *testing.T) {
	var text string = "My name is chandra sekhar. I'm learning GO language."
	wordCount := wordCount(text)
	want := 9
	if wordCount != want {
		t.Errorf("wordCount = %d; want %d", wordCount, want)
	}
}

func TestCharCount(t *testing.T) {
	var text string = "Hello"
	charCount := charCount(text)
	want := 5
	if charCount != want {
		t.Errorf("charCount = %d; want %d", charCount, want)
	}
}

func TestLineCount(t *testing.T) {
	var text string = "Hello\n World"
	lineCount := lineCount(text)
	want := 2
	if lineCount != want {
		t.Errorf("lineCount = %d; want %d", lineCount, want)
	}
}
