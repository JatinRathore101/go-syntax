package main

import "fmt"

// Short variable declaration := (most common in Go, only inside functions)
func main() {
	// := infers type and declares in one step
	city := "Bangalore"
	fmt.Println("Short declaration:", city)

	// Works only inside functions — not at package level
	isActive := true
	count := 42
	fmt.Println("isActive:", isActive, "| count:", count)

	// Reassign using = (not :=)
	count = 100
	fmt.Println("Reassigned count:", count)
}
