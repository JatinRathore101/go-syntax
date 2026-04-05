package main

import "fmt"

// Creating slices and accessing elements
func main() {
	// Slice literal (dynamic size, unlike arrays)
	numbers := []int{10, 20, 30}

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	// Access by index
	fmt.Println("First:", numbers[0])
	fmt.Println("Second:", numbers[1])

	// Update element
	numbers[0] = 99
	fmt.Println("After update:", numbers)
}
