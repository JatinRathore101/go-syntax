package main

import "fmt"

// Boolean and String types in Go
func main() {
	// Boolean: true or false
	var isActive bool = true
	var isLoggedIn bool // zero value is false

	fmt.Println("isActive:", isActive, "| Type:", fmt.Sprintf("%T", isActive))
	fmt.Println("isLoggedIn (zero value):", isLoggedIn)

	// String: sequence of characters, enclosed in double quotes
	var name string = "Jai"
	greeting := "Hello, " + name // string concatenation

	fmt.Println("name:", name, "| Type:", fmt.Sprintf("%T", name))
	fmt.Println("greeting:", greeting)

	// String length (in bytes)
	fmt.Println("Length of name:", len(name))
}
