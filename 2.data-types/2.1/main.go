package main

import "fmt"

// Integer and Float types in Go
func main() {
	// Integer types: int, int8, int16, int32, int64
	var a int = 10
	var b int8 = 120
	var c int64 = 9000000000

	fmt.Println("int:", a, "| Type:", fmt.Sprintf("%T", a))
	fmt.Println("int8:", b, "| Type:", fmt.Sprintf("%T", b))
	fmt.Println("int64:", c, "| Type:", fmt.Sprintf("%T", c))

	// Float types: float32, float64
	var price float32 = 19.99
	var pi float64 = 3.1415926535

	fmt.Println("float32:", price, "| Type:", fmt.Sprintf("%T", price))
	fmt.Println("float64:", pi, "| Type:", fmt.Sprintf("%T", pi))
}
