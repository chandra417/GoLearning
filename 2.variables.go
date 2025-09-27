package main

import "fmt"

func main() {
	// Declare a variable with explicit type
	var message string = "Welcome to Go!"
	fmt.Println(message)

	// Short variable declaration (type inferred)
	number := 42
	fmt.Println(number)

	// Multiple types
	var pi float64 = 3.1415
	var isReady bool = true

	fmt.Println(pi, isReady)
}
