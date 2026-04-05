package main

import (
	"fmt"
	"sort"
	"strings"
)

// Split, join, and custom sorting of strings
func main() {
	// Split a string into a slice
	parts := strings.Split("a,b,c", ",")
	fmt.Println("Split:", parts)

	// Join a slice into a string
	joined := strings.Join(parts, "-")
	fmt.Println("Join:", joined)

	// Case-insensitive sort
	words := []string{"Go", "java", "Python", "c"}
	sort.Slice(words, func(i, j int) bool {
		return strings.ToLower(words[i]) < strings.ToLower(words[j])
	})
	fmt.Println("Case-insensitive sort:", words)

	// Sort by string length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	fmt.Println("Sort by length:", words)
}
