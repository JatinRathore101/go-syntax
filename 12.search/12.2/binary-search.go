package main

import (
	"fmt"
	"sort"
)

// Binary search: O(log n), requires a sorted array
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11}

	// Manual binary search
	fmt.Println("Binary search for 7:", binarySearch(arr, 7))
	fmt.Println("Binary search for 4:", binarySearch(arr, 4))

	// Built-in binary search from sort package
	index := sort.SearchInts(arr, 7)
	if index < len(arr) && arr[index] == 7 {
		fmt.Println("sort.SearchInts found 7 at index:", index)
	} else {
		fmt.Println("sort.SearchInts: not found")
	}
}
