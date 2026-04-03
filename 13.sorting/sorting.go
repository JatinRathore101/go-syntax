package main

import (
	"fmt"
	"sort"
)

func main() {

	// -----------------------------------------
	// 1. SORT INTEGERS
	// -----------------------------------------
	nums := []int{5, 2, 8, 1, 3}

	sort.Ints(nums) // ascending order
	fmt.Println("Sorted ints:", nums)

	// Descending (common trick)
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	fmt.Println("Sorted ints (desc):", nums)

	// -----------------------------------------
	// 2. SORT STRINGS
	// -----------------------------------------
	strs := []string{"banana", "apple", "cherry"}

	sort.Strings(strs) // alphabetical
	fmt.Println("Sorted strings:", strs)

	// -----------------------------------------
	// 3. CUSTOM SORT (IMPORTANT)
	// -----------------------------------------
	// sort.Slice lets you define your own comparator

	arr := []int{10, 3, 25, 7}

	// Example: sort based on last digit
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]%10 < arr[j]%10
	})
	fmt.Println("Custom sort (last digit):", arr)

	// -----------------------------------------
	// 4. SORT STRUCTS (VERY COMMON)
	// -----------------------------------------
	people := []Person{
		{"Alice", 25},
		{"Bob", 20},
		{"Charlie", 30},
	}

	// Sort by Age (ascending)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sorted by Age:", people)

	// Sort by Name (alphabetical)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("Sorted by Name:", people)

	// -----------------------------------------
	// 5. STABLE SORT (preserves order of equals)
	// -----------------------------------------
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Stable sort by Age:", people)
}

// -----------------------------------------
// STRUCT DEFINITION
// -----------------------------------------
type Person struct {
	Name string
	Age  int
}
