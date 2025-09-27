package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func main() {
	addResult := add(10, 20)
	fmt.Println(addResult)

	multiplyResult := multiply(10, 20)
	fmt.Println(multiplyResult)
}
