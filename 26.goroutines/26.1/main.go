package main

import (
	"fmt"
	"time"
)

// Goroutine basics: run functions concurrently using the 'go' keyword
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting goroutine...")

	// 'go' keyword starts the function in a new goroutine (concurrent)
	go printNumbers()

	// Main must stay alive long enough for the goroutine to finish
	// Without this sleep, main exits before printNumbers completes
	time.Sleep(2 * time.Second)

	fmt.Println("Done")
}
