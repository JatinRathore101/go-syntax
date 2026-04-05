package main

import "fmt"

// Multiple return values and named return values
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// Named return values: variables declared in the signature
func rectangleStats(length int, width int) (area int, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // bare return uses named variables
}

func main() {
	// Multiple return values
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)

	// Named return values
	area, perimeter := rectangleStats(10, 5)
	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)
}
