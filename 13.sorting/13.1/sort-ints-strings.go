package main

import (
	"fmt"
	"sort"
)

// Sorting integers and strings using the sort package
func main() {
	// Sort integers ascending
	nums := []int{5, 2, 8, 1, 3}
	sort.Ints(nums)
	fmt.Println("Sorted ints (asc):", nums)

	// Sort integers descending using sort.Slice
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println("Sorted ints (desc):", nums)

	// Sort strings alphabetically
	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println("Sorted strings:", strs)
}
