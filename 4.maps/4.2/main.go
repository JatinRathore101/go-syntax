package main

import "fmt"

// Add, update, check existence, and delete map entries
func main() {
	scores := map[string]int{"math": 90, "science": 85}

	// Add a new key
	scores["english"] = 88
	fmt.Println("After add:", scores)

	// Update an existing key
	scores["math"] = 95
	fmt.Println("After update:", scores)

	// Check if a key exists (two-value assignment)
	value, exists := scores["science"]
	if exists {
		fmt.Println("Science exists, value:", value)
	} else {
		fmt.Println("Science not found")
	}

	// Delete a key
	delete(scores, "english")
	fmt.Println("After delete:", scores)
}
