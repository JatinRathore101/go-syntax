package main

import "fmt"

// Appending elements and joining two slices
func main() {
	numbers := []int{10, 20, 30}

	// Append a single element
	numbers = append(numbers, 40)
	fmt.Println("After append:", numbers)

	// Append multiple elements
	numbers = append(numbers, 50, 60)
	fmt.Println("After multi-append:", numbers)

	// Join two slices using ... (spread operator)
	more := []int{70, 80, 90}
	numbers = append(numbers, more...)
	fmt.Println("After joining:", numbers)
}
