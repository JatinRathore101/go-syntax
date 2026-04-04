package main

import "fmt"

// Basic for loop and while-style loop (Go has only 'for', no 'while')
func main() {
	// Classic for loop: init; condition; post
	fmt.Println("Basic for loop:")
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	// While-style: only condition, no init or post
	fmt.Println("\nWhile-style loop:")
	j := 0
	for j < 5 {
		fmt.Println("j =", j)
		j++
	}

	// Infinite loop with break
	fmt.Println("\nInfinite loop (break at 3):")
	k := 0
	for {
		fmt.Println("k =", k)
		k++
		if k == 3 {
			break
		}
	}
}
