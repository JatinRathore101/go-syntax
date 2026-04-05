package main

import "fmt"

// Sub-slicing and iterating over a slice with range
func main() {
	numbers := []int{10, 20, 30, 40, 50, 60}

	// Sub-slice: numbers[start:end] (end is exclusive)
	firstThree := numbers[0:3]
	fmt.Println("First three:", firstThree)

	lastTwo := numbers[4:]
	fmt.Println("Last two:", lastTwo)

	// Iterate with index and value
	fmt.Println("All elements:")
	for index, value := range numbers {
		fmt.Println(" Index:", index, "Value:", value)
	}

	// Iterate with value only (ignore index with _)
	fmt.Println("Values only:")
	for _, value := range numbers {
		fmt.Println(" ", value)
	}
}
