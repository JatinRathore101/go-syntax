package main

import "fmt"

// Passing a pointer to a function to modify the original value
func updateValue(x *int) {
	*x = 100 // dereference and assign
}

func main() {
	x := 10
	fmt.Println("Before:", x)

	updateValue(&x) // pass address of x

	fmt.Println("After updateValue:", x) // 100
}
