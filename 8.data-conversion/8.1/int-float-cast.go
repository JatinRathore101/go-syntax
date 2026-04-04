package main

import "fmt"

// Basic numeric type casting: int ↔ float
func main() {
	// int → float64
	var a int = 10
	b := float64(a)
	fmt.Println("int:", a, "-> float64:", b)

	// float64 → int (decimal part is truncated, not rounded)
	var x float64 = 9.8
	y := int(x)
	fmt.Println("float64:", x, "-> int:", y)

	// int32 → int64
	var small int32 = 100
	big := int64(small)
	fmt.Println("int32:", small, "-> int64:", big)
}
