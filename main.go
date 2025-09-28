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

	// VARIABLE: Integer
	var a int = 10 // Standard integer
	var b int64 = 12345678912345
	var c uint = 20 // Unsigned integer (no negatives)
	fmt.Println(a, b, c)

	// VARIABLE: Floating	-	Use float32 or float64 for decimals (64-bit is default).
	var pi float64 = 3.14159
	area := 5.5 * 7.4 // float64 inferred
	fmt.Println(pi, area)

	// VARIABLE: Boolean	-	Booleans are true or false.
	var isReady bool = true
	done := false
	fmt.Println(isReady, done)

	// VARIABLE: string	-	Strings are sequences of characters.
	var name string = "Chandrasekhar"
	msg := "Welcome to Go!"
	fmt.Println(name, msg)

	// VARIABLE: Array	-	Arrays have a fixed length defined at creation.
	var arr [3]int = [3]int{10, 20, 30} // Fixed size
	fmt.Println(arr)

	// VARIABLE: Slice	-	Slices are dynamic arrays (most used for lists).
	nums := []int{1, 2, 3, 4, 5} // Flexible, like an arraylist
	nums = append(nums, 6)       // Add more items
	fmt.Println(nums)

	// VARIABLE: Map	-	Maps are key-value stores.
	capitals := map[string]string{
		"India":   "New Delhi",
		"France":  "Paris",
		"Germany": "Berlin",
	}

	fmt.Println("Capital of India:", capitals["India"])

	capitals["Japan"] = "Tokyo"
	fmt.Println("Added Japan:", capitals["Japan"])

	capitals["Germany"] = "Munich"
	fmt.Println("Updated Germany:", capitals["Germany"])

	delete(capitals, "France")

	if val, exists := capitals["France"]; exists {
		fmt.Println("Found France:", val)
	} else {
		fmt.Println("No France")
	}

	for country, capital := range capitals {
		fmt.Printf("Country: %s, Capital: %s\n", country, capital)
	}

	// VARIABLE: Struct
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.Name, p.Age)

}
