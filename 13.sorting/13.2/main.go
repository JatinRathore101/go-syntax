package main

import (
	"fmt"
	"sort"
)

// Custom sort using sort.Slice with a comparator function
func main() {
	// Sort by last digit
	arr := []int{10, 3, 25, 7}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i]%10 < arr[j]%10
	})
	fmt.Println("Sort by last digit:", arr)

	// Sort strings by length
	words := []string{"Go", "Python", "C", "Java"}
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	fmt.Println("Sort by length:", words)
}
