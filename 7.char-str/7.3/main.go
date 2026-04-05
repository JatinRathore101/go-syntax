package main

import (
	"fmt"
	"strings"
)

// String trimming and searching
func main() {
	// Trim whitespace
	str := "   hello world   "
	fmt.Println("TrimSpace:", strings.TrimSpace(str))
	fmt.Println("TrimLeft:", strings.TrimLeft(str, " "))
	fmt.Println("TrimRight:", strings.TrimRight(str, " "))

	// Trim specific characters
	fmt.Println("Trim '!':", strings.Trim("!!hello!!", "!"))

	// Search within a string
	text := "golang is awesome"
	fmt.Println("Contains 'go':", strings.Contains(text, "go"))
	fmt.Println("HasPrefix 'go':", strings.HasPrefix(text, "go"))
	fmt.Println("HasSuffix 'me':", strings.HasSuffix(text, "me"))
	fmt.Println("Index of 'is':", strings.Index(text, "is"))
}
