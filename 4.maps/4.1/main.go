package main

import "fmt"

// Creating a map and accessing values
func main() {
	// Create a map: map[KeyType]ValueType
	scores := map[string]int{
		"math":    90,
		"science": 85,
	}
	fmt.Println("Map:", scores)

	// Access a value by key
	fmt.Println("Math score:", scores["math"])

	// Accessing a missing key returns the zero value (0 for int)
	fmt.Println("History (missing key):", scores["history"])
}
