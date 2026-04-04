package main

import "fmt"

// Basic pointers: address, dereference, and modify via pointer
func main() {
	a := 10

	// & gives the memory address of a variable
	var p *int = &a
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a (&a):", &a)
	fmt.Println("Pointer p (stores address):", p)
	fmt.Println("Value at pointer (*p):", *p) // dereference

	// Modify original value through pointer
	*p = 20
	fmt.Println("After *p = 20, a is:", a)

	// Short pointer syntax
	b := 50
	pb := &b
	fmt.Println("b:", b, "| *pb:", *pb)
}
