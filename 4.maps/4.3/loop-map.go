package main

import "fmt"

// Looping over a map using range
func main() {
	scores := map[string]int{
		"math":    95,
		"science": 88,
		"english": 76,
	}

	// Loop over key-value pairs
	fmt.Println("All scores:")
	for key, value := range scores {
		fmt.Println(" ", key, "->", value)
	}

	// Loop over keys only
	fmt.Println("Keys only:")
	for key := range scores {
		fmt.Println(" ", key)
	}
}
