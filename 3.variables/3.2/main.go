package main

import "fmt"

// var declarations: explicit type and type inference
func main() {
	// Explicit type declaration
	var age int = 25
	fmt.Println("Explicit type (int):", age)

	// Type inference: Go detects the type automatically
	var name = "Jai"
	fmt.Println("Type inferred (string):", name)

	// Go does NOT have 'let' or 'const' for local mutable vars — only var and :=
	var score float64 = 9.5
	fmt.Println("Float var:", score)
}
