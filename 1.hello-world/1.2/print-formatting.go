package main

import "fmt"

// Printing different types of values using fmt.Println and fmt.Printf
func main() {
	// Print a plain string
	fmt.Println("Hello, Go!")

	// Print with formatting: %s = string, %d = int, %f = float
	fmt.Printf("Name: %s, Age: %d, Score: %.2f\n", "Jai", 25, 98.5)

	// Print multiple values
	fmt.Println("Sum:", 10+20)
}
