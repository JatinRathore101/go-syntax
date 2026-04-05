package main

import (
	"fmt"
	"strings"
)

// Case conversion: ToUpper and ToLower
func main() {
	s := "Hello GoLang 123"

	fmt.Println("Original:", s)
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("ToUpper:", strings.ToUpper(s))
}
