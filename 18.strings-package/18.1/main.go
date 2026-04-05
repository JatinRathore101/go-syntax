package main

import (
	"fmt"
	"strings"
)

// Contains, HasPrefix, HasSuffix: check string content
func main() {
	text := "hello golang world"

	fmt.Println("Contains 'golang':", strings.Contains(text, "golang"))
	fmt.Println("HasPrefix 'hello':", strings.HasPrefix(text, "hello"))
	fmt.Println("HasSuffix 'world':", strings.HasSuffix(text, "world"))
}
