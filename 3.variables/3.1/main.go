package main

import "fmt"

// Global variables and constants: accessible throughout the file
var globalVar string = "I am a global variable"

// Constants cannot be changed after declaration
const pi float64 = 3.14159

func main() {
	fmt.Println("Global variable:", globalVar)
	fmt.Println("Constant pi:", pi)

	// Constants can be used in expressions
	radius := 5.0
	area := pi * radius * radius
	fmt.Printf("Area of circle (r=5): %.2f\n", area)
}
