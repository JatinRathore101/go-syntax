package main

import (
	"fmt"
	"time"
)

// ------------------------------
// What are Goroutines?
// ------------------------------
// Goroutines are lightweight threads managed by Go runtime.
// They allow you to run functions concurrently (at the same time).
// Extremely cheap compared to OS threads (~2KB stack vs ~1MB stack for threads).

// ------------------------------
// Why Use Goroutines?
// ------------------------------
// 1. Concurrency: Run multiple tasks at once.
// 2. Efficiency: Better CPU utilization.
// 3. Non-blocking: Prevent one task from stopping others.
// 4. Suitable for IO-bound tasks, network requests, or parallel computations.

// ------------------------------
// Example Function
// ------------------------------
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond) // Simulate work
	}
}

func printLetters() {
	for _, l := range "ABCDE" {
		fmt.Println("Letter:", string(l))
		time.Sleep(700 * time.Millisecond) // Simulate work
	}
}

func main() {
	// ------------------------------
	// Running Goroutines
	// ------------------------------
	fmt.Println("Starting Goroutines...")

	// Use "go" keyword to start a goroutine
	go printNumbers() // runs concurrently
	go printLetters() // runs concurrently

	// Without sleep or sync, main may exit before goroutines finish
	time.Sleep(5 * time.Second) // wait enough time for goroutines to finish

	fmt.Println("Goroutines finished.")
}

// ------------------------------
// Key Points:
// ------------------------------
// 1. Use `go` keyword before function call to start a goroutine.
// 2. Main function must stay alive; otherwise goroutines may not complete.
// 3. Use goroutines for concurrent tasks like:
//    - Network calls / HTTP requests
//    - Reading/writing files concurrently
//    - Parallel computations
//    - Background tasks
// 4. Advantage: Lightweight, fast, and easy to use concurrency.
// 5. Don't use goroutines if tasks must run strictly in order (then normal functions are enough).
