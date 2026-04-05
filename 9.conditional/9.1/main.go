package main

import "fmt"

// if, if-else, and if-else-if ladder
func main() {
	// Simple if
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	}

	// if-else
	score := 40
	if score >= 50 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	// if-else-if ladder (Go uses 'else if', not 'elif')
	marks := 82
	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 75 {
		fmt.Println("Grade: B")
	} else if marks >= 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: Fail")
	}
}
