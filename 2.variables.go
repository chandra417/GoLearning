package main

import "fmt"

func main() {
	// Declare a variable with explicit type
	var name string = "Chandrasekhar"
	fmt.Println(name)

	// Short variable declaration (type inferred)
	my_number := 17
	fmt.Println(my_number)

	// Multiple types
	var pi float64 = 3.1415
	var isReady bool = true

	fmt.Println(pi, isReady)
}
