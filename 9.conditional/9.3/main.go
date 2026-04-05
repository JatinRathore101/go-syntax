package main

import "fmt"

// Switch without a condition: acts like a cleaner if-else ladder
func main() {
	temp := 28

	// No expression after switch — each case is a boolean condition
	switch {
	case temp > 35:
		fmt.Println("Very Hot")
	case temp > 25:
		fmt.Println("Warm")
	case temp > 15:
		fmt.Println("Cool")
	default:
		fmt.Println("Cold")
	}
}
