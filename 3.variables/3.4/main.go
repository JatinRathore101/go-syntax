package main

import "fmt"

// Declaring multiple variables at once
func main() {
	// Multiple variables on one line
	var a, b int = 10, 20
	fmt.Println("Multiple var:", a, b)

	// Multiple short declarations
	x, y := 100, 200
	fmt.Println("Multiple :=:", x, y)

	// Block-style declaration (grouped)
	var (
		username string = "admin"
		password string = "secret"
		port     int    = 8080
	)
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)
	fmt.Println("Port:", port)
}
