package main

import (
	"fmt"
	"sort"
)

// Sorting a slice of structs by field, and stable sort
type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
	}

	// Sort by Age ascending
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by Age:", people)

	// Sort by Name alphabetically
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("Sorted by Name:", people)

	// Stable sort: preserves original order for equal elements
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Stable sort by Age:", people)
}
