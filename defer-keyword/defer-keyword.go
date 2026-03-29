package main

import (
	"fmt"
	"os"
)

func main() {

	// -----------------------------------------
	// 1. BASIC USAGE OF defer
	// -----------------------------------------
	// defer delays execution until the surrounding function returns

	defer fmt.Println("This runs LAST")
	fmt.Println("This runs FIRST")

	// Output order:
	// This runs FIRST
	// This runs LAST

	// -----------------------------------------
	// 2. MULTIPLE defer (LIFO order)
	// -----------------------------------------
	// Last deferred = executed first (stack behavior)

	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")

	// Execution order at end:
	// defer 3
	// defer 2
	// defer 1

	// -----------------------------------------
	// 3. COMMON REAL USE CASE: closing resources
	// -----------------------------------------
	// Always defer closing files, DB, etc.

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close() // ensures file is closed no matter what

	fmt.Println("File opened successfully")

	// -----------------------------------------
	// 4. ARGUMENTS ARE EVALUATED IMMEDIATELY
	// -----------------------------------------
	// Values are captured at the time of defer, not execution time

	x := 10
	defer fmt.Println("Deferred value:", x) // x = 10 now
	x = 20

	// Output will be: Deferred value: 10

	// -----------------------------------------
	// 5. DEFER WITH FUNCTIONS (common pattern)
	// -----------------------------------------
	defer func() {
		fmt.Println("Deferred anonymous function")
	}()

	// -----------------------------------------
	// 6. MODIFY NAMED RETURN VALUES (advanced but useful)
	// -----------------------------------------
	fmt.Println("Result:", add(5, 3))
}

// -----------------------------------------
// Function showing defer modifying return
// -----------------------------------------
func add(a, b int) (result int) {
	result = a + b

	// This runs AFTER return statement but BEFORE actual return
	defer func() {
		result += 1 // modifies return value
	}()

	return // returns (a + b + 1)
}
