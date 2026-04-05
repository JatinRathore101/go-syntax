package main

import "fmt"

// Basic functions: no params, with params, with return value
func greet() {
	fmt.Println("Hello! Welcome to Go functions.")
}

func greetUser(name string) {
	fmt.Println("Hello,", name)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	greet()

	greetUser("Jai")

	result := add(5, 3)
	fmt.Println("5 + 3 =", result)
}
