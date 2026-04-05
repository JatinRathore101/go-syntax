package main

import "fmt"

// defer argument evaluation and modifying named return values
func main() {
	// Arguments are evaluated at the time of defer, not at execution
	x := 10
	defer fmt.Println("Deferred x (captured as 10):", x) // captures x=10 now
	x = 20
	fmt.Println("Current x:", x) // 20

	// defer with anonymous function (captures by reference)
	y := 10
	defer func() {
		fmt.Println("Anonymous defer sees y:", y) // sees y=20
	}()
	y = 20

	// Named return value modified by defer
	fmt.Println("add(5,3) =", add(5, 3)) // returns 9 (5+3+1)
}

// defer can modify named return values
func add(a, b int) (result int) {
	result = a + b
	defer func() {
		result += 1 // modifies the return value before it's returned
	}()
	return
}
