package main

import (
	"fmt"
	"strings"
)

// Split, Join, and Replace
func main() {
	// Split a string into a slice
	str := "apple,banana,mango"
	parts := strings.Split(str, ",")
	fmt.Println("Split:", parts)

	// Join a slice into a string
	joined := strings.Join(parts, " | ")
	fmt.Println("Joined:", joined)

	// Replace: replace first N occurrences (-1 = all)
	sentence := "I like java and java is great"
	fmt.Println("Replace 1:", strings.Replace(sentence, "java", "go", 1))
	fmt.Println("Replace all:", strings.ReplaceAll(sentence, "java", "go"))
}
