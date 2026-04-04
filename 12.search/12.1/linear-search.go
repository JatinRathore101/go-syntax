package main

import "fmt"

// Linear search: O(n), works on any unsorted array
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i // return index if found
		}
	}
	return -1 // not found
}

func main() {
	arr := []int{5, 3, 8, 1, 9, 2}

	idx := linearSearch(arr, 8)
	if idx != -1 {
		fmt.Println("Found 8 at index:", idx)
	} else {
		fmt.Println("8 not found")
	}

	fmt.Println("Search for 99:", linearSearch(arr, 99)) // -1
}
