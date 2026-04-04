package main

import "fmt"

// Variadic functions: accept any number of arguments
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func main() {
	// Pass any number of arguments
	fmt.Println("Sum(1,2,3):", sum(1, 2, 3))
	fmt.Println("Sum(1..5):", sum(1, 2, 3, 4, 5))

	// Spread a slice into a variadic function using ...
	nums := []int{10, 20, 30}
	fmt.Println("Sum from slice:", sum(nums...))
}
