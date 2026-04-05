package main

import (
	"fmt"
	"unicode"
)

// Checking character types using the unicode package
func main() {
	ch := 'A' // rune (a single character)

	fmt.Println("IsLetter:", unicode.IsLetter(ch))
	fmt.Println("IsDigit:", unicode.IsDigit(ch))
	fmt.Println("IsSpace:", unicode.IsSpace(ch))
	fmt.Println("IsAlphanumeric:", unicode.IsLetter(ch) || unicode.IsDigit(ch))

	// Loop over a string and classify each character
	fmt.Println("\nCharacter types in \"Go123!\":")
	for _, r := range "Go123!" {
		switch {
		case unicode.IsLetter(r):
			fmt.Println(string(r), "-> Letter")
		case unicode.IsDigit(r):
			fmt.Println(string(r), "-> Digit")
		default:
			fmt.Println(string(r), "-> Other")
		}
	}
}
