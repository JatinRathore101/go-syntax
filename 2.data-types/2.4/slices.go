package main

import "fmt"

// Slices in Go: dynamic-size lists (most used collection type)
func main() {
	// Create a slice (no fixed size, unlike arrays)
	scores := []int{90, 85, 70}

	fmt.Println("Slice:", scores)
	fmt.Println("Type:", fmt.Sprintf("%T", scores))
	fmt.Println("Length:", len(scores), "| Capacity:", cap(scores))

	// Append element to slice
	scores = append(scores, 95)
	fmt.Println("After append:", scores)

	// Access by index
	fmt.Println("First score:", scores[0])
}
