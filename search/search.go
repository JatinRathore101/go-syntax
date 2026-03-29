package main

import (
	"fmt"
	"sort"
)

func main() {

	// Sample sorted array (IMPORTANT for binary search)
	arr := []int{1, 3, 5, 7, 9, 11}

	// -----------------------------------------
	// 1. LINEAR SEARCH (works on any array)
	// -----------------------------------------
	fmt.Println("Linear Search (find 7):", linearSearch(arr, 7))

	// -----------------------------------------
	// 2. BINARY SEARCH (manual implementation)
	// -----------------------------------------
	// Works ONLY on sorted arrays

	fmt.Println("Binary Search (find 7):", binarySearch(arr, 7))

	// -----------------------------------------
	// 3. BINARY SEARCH (built-in from sort pkg)
	// -----------------------------------------
	// Returns index where element is found OR where it should be inserted

	index := sort.SearchInts(arr, 7)

	if index < len(arr) && arr[index] == 7 {
		fmt.Println("Built-in Binary Search: Found at index", index)
	} else {
		fmt.Println("Built-in Binary Search: Not found")
	}
}

// -----------------------------------------
// LINEAR SEARCH
// Time: O(n)
// -----------------------------------------
func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i // return index if found
		}
	}
	return -1 // not found
}

// -----------------------------------------
// BINARY SEARCH (ITERATIVE)
// Time: O(log n)
// -----------------------------------------
func binarySearch(arr []int, target int) int {

	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid // found
		} else if arr[mid] < target {
			left = mid + 1 // search right half
		} else {
			right = mid - 1 // search left half
		}
	}

	return -1 // not found
}
