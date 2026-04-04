package main

import "fmt"

// Arrays in Go: fixed-size, same-type collections
func main() {
	// Declare and initialize an array of 3 ints
	var numbers [3]int = [3]int{10, 20, 30}

	fmt.Println("Array:", numbers)
	fmt.Println("Type:", fmt.Sprintf("%T", numbers))
	fmt.Println("Length:", len(numbers))

	// Access by index
	fmt.Println("First element:", numbers[0])

	// Update element
	numbers[1] = 99
	fmt.Println("After update:", numbers)

	// Short declaration
	fruits := [3]string{"apple", "banana", "cherry"}
	fmt.Println("Fruits:", fruits)
}
