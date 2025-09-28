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

func charCount(text string) int {
	return len(text)
}

func lineCount(text string) int {
	return len(strings.Split(text, "\n"))
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
	fmt.Print("			// VARIABLE: Floating - Use float32 or float64 for decimals (64-bit is default).\n")
	fmt.Print("====================================================\n")
	var pi float64 = 3.14159
	area := 5.5 * 7.4 // float64 inferred
	fmt.Println(pi, area)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Boolean - Booleans are true or false.\n")
	fmt.Print("====================================================\n")
	var isReady bool = true
	done := false
	fmt.Println(isReady, done)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: string - Strings are sequences of characters.\n")
	fmt.Print("====================================================\n")
	var name string = "Chandrasekhar"
	msg := "Welcome to Go!"
	fmt.Println(name, msg)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Array - Arrays have a fixed length defined at creation.\n")
	fmt.Print("====================================================\n")

	var arr = [3]int{10, 20, 30} // Fixed size
	var arrayNums [5]int         //This creates an array of 5 integers, all initialized to zero.
	var arrayNames = [3]string{"A", "B", "C"}
	var arrayActualNums = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)
	fmt.Println(arrayNums)
	fmt.Println(arrayNames)
	fmt.Println(arrayActualNums)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Slice - Slices are dynamic arrays (most used for lists).\n")
	fmt.Print("====================================================\n")
	nums := []int{1, 2, 3, 4, 5} // Flexible, like an arraylist
	nums = append(nums, 6)       // Add more items
	fmt.Println(nums)

	var scores []int
	scores = []int{30, 40, 20, 33}
	fmt.Println(scores)

	fmt.Println(scores[1:3])

	// Few examples
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(s))
	fmt.Println(cap(s))

	s = append(s, 10, 11)
	fmt.Println(s)

	test := make([]int, 2)
	new_test := []int{100, 200}

	copy(test, new_test)
	fmt.Println(test)

	fmt.Print("====================================================\n")
	fmt.Print("			// VARIABLE: Map - Maps are key-value stores.\n")
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
	fmt.Print("			// VARIABLE: Struct - Used to define custom data structures—like classes without methods.\n")
	fmt.Print("====================================================\n")
	type Person struct {
		Name string
		Age  int
	}

	p1 := Person{Name: "rob", Age: 30}
	p2 := Person{Name: "bob", Age: 20}
	p3 := Person{Name: "pop", Age: 10}
	fmt.Println(p1.Name, p1.Age)
	fmt.Println(p2.Name, p2.Age)
	fmt.Println(p3.Name, p3.Age)

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
		os.Exit(1)
		return
	}

	text := string(contents)
	wordCount := wordCount(text)
	fmt.Printf("Word Count: %d\n", wordCount)

	charCount := charCount(text)
	fmt.Printf("Char Count: %d\n", charCount)

	lineCount := lineCount(text)
	fmt.Printf("Line Count: %d\n", lineCount)

	fmt.Print("====================================================\n")
}
