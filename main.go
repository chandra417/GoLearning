package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func oddEven(a int) string {
	if a%2 == 0 {
		return fmt.Sprintf("%d is a even number\n", a)
	} else {
		return fmt.Sprintf("%d is a odd number\n", a)
	}
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is my first commit")
}
