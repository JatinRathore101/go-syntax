package main

import (
	"fmt"
	"strings"
)

// ToUpper, ToLower, TrimSpace, Count, Index
func main() {
	// Case conversion
	name := "GoLang"
	fmt.Println("Upper:", strings.ToUpper(name))
	fmt.Println("Lower:", strings.ToLower(name))

	// Trim whitespace
	raw := "   hello world   "
	fmt.Println("TrimSpace: [" + strings.TrimSpace(raw) + "]")

	// Count occurrences
	data := "go go go"
	fmt.Println("Count 'go':", strings.Count(data, "go"))

	// Index of first occurrence (-1 if not found)
	fmt.Println("Index of 'go':", strings.Index(data, "go"))
}
