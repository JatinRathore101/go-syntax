package main

import "fmt"

// break and continue in loops
func main() {
	// continue: skip current iteration
	fmt.Println("Continue (skip i=2):")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("i =", i)
	}

	// break: exit the loop early
	fmt.Println("\nBreak (stop at i=3):")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("i =", i)
	}
}
