package main

import "fmt"

// Multiple defers execute in LIFO (last in, first out) order
func main() {
	defer fmt.Println("defer 1 (runs 3rd)")
	defer fmt.Println("defer 2 (runs 2nd)")
	defer fmt.Println("defer 3 (runs 1st)")

	fmt.Println("Normal execution")

	// Output order:
	// Normal execution
	// defer 3 (runs 1st)
	// defer 2 (runs 2nd)
	// defer 1 (runs 3rd)
}
