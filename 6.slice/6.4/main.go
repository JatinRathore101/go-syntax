package main

import "fmt"

// Pop, remove by index, copy, and reference behavior of slices
func main() {
	numbers := []int{10, 20, 30, 40, 50}

	// Pop last element (Go has no built-in pop)
	numbers = numbers[:len(numbers)-1]
	fmt.Println("After pop:", numbers)

	// Remove element at index 1 (splice pattern)
	i := 1
	numbers = append(numbers[:i], numbers[i+1:]...)
	fmt.Println("After removing index 1:", numbers)

	// Copy a slice (independent copy)
	original := []int{1, 2, 3}
	copySlice := append([]int{}, original...)
	copySlice[0] = 999
	fmt.Println("Original:", original)   // unchanged
	fmt.Println("Copy:", copySlice)

	// Slices are reference types: assignment shares the same backing array
	a := []int{100, 200, 300}
	b := a
	b[0] = 999
	fmt.Println("a (affected):", a) // a is also changed
	fmt.Println("b:", b)
}
