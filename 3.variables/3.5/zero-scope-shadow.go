package main

import "fmt"

// Zero values, variable scope, and shadowing
func main() {
	// Zero values: Go assigns defaults when not initialized
	var num int     // 0
	var text string // ""
	var flag bool   // false

	fmt.Println("Zero int:", num)
	fmt.Println("Zero string:", "["+text+"]")
	fmt.Println("Zero bool:", flag)

	// Scope: variables exist only within their block
	outer := "I am in main()"
	fmt.Println("Outer:", outer)

	{
		inner := "I exist only in this block"
		fmt.Println("Inner:", inner)
		// outer is accessible here too
	}
	// inner is NOT accessible here

	// Shadowing: inner block variable overrides outer
	value := 10
	fmt.Println("Before shadow:", value)
	{
		value := 20 // new variable, shadows outer
		fmt.Println("Shadowed value:", value)
	}
	fmt.Println("Outer value unchanged:", value)
}
