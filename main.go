package main

import (
	"fmt"
	"os"
	"strings"
)

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func oddNumber(a int) bool {
	return a%2 != 0
}

func wordCount(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is my first commit")

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Integer\n")
	fmt.Print("====================================================\n")
	var a int = 10 // Standard integer
	var b int64 = 12345678912345
	var c uint = 20 // Unsigned integer (no negatives)
	fmt.Println(a, b, c)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Floating\t-\tUse float32 or float64 for decimals (64-bit is default).\n")
	fmt.Print("====================================================\n")
	var pi float64 = 3.14159
	area := 5.5 * 7.4 // float64 inferred
	fmt.Println(pi, area)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Boolean\t-\tBooleans are true or false.\n")
	fmt.Print("====================================================\n")
	var isReady bool = true
	done := false
	fmt.Println(isReady, done)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: string\t-\tStrings are sequences of characters.\n")
	fmt.Print("====================================================\n")
	var name string = "Chandrasekhar"
	msg := "Welcome to Go!"
	fmt.Println(name, msg)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Array\t-\tArrays have a fixed length defined at creation.\n")
	fmt.Print("====================================================\n")
	var arr [3]int = [3]int{10, 20, 30} // Fixed size
	fmt.Println(arr)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Slice\t-\tSlices are dynamic arrays (most used for lists).\n")
	fmt.Print("====================================================\n")
	nums := []int{1, 2, 3, 4, 5} // Flexible, like an arraylist
	nums = append(nums, 6)       // Add more items
	fmt.Println(nums)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Map\t-\tMaps are key-value stores.\n")
	fmt.Print("====================================================\n")
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

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Struct\n")
	fmt.Print("====================================================\n")
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "Alice", Age: 30}
	fmt.Println(p.Name, p.Age)

	fmt.Print("====================================================\n")
	fmt.Print("			// CLI word count from a file\n")
	fmt.Print("====================================================\n")
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage: go run main.go filename.txt")
		return
	}
	filename := os.Args[1]
	contents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %s\n", filename, err)
		return
	}

	text := string(contents)
	wordCount := wordCount(text)
	fmt.Printf("Word Count: %d\n", wordCount)
	fmt.Print("====================================================\n")
}
