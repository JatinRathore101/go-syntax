package main

import "fmt"

// Pass by value vs pass by pointer, anonymous functions, and function variables
func changeValue(x int) {
	x = 100 // only changes local copy
	fmt.Println("Inside changeValue:", x)
}

func changeValuePointer(x *int) {
	*x = 100 // modifies original
}

func main() {
	// Pass by value: original is NOT changed
	num := 10
	fmt.Println("Before changeValue:", num)
	changeValue(num)
	fmt.Println("After changeValue:", num) // still 10

	// Pass by pointer: original IS changed
	changeValuePointer(&num)
	fmt.Println("After changeValuePointer:", num) // 100

	// Anonymous function: defined and called immediately
	func() {
		fmt.Println("Anonymous function called")
	}()

	// Function stored in a variable
	multiply := func(a int, b int) int {
		return a * b
	}
	fmt.Println("4 * 5 =", multiply(4, 5))
}
