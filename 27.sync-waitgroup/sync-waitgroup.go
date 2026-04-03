package main

import (
	"fmt"
	"sync"
	"time"
)

// ------------------------------
// What is WaitGroup?
// ------------------------------
// WaitGroup is a synchronization primitive that waits for a collection
// of goroutines to finish.
// It is part of "sync" package.

// ------------------------------
// Example Functions
// ------------------------------
func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done() // signals when goroutine is finished
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(300 * time.Millisecond) // simulate work
	}
}

func printLetters(wg *sync.WaitGroup) {
	defer wg.Done() // signals when goroutine is finished
	for _, l := range "ABCDE" {
		fmt.Println("Letter:", string(l))
		time.Sleep(500 * time.Millisecond) // simulate work
	}
}

func main() {
	// ------------------------------
	// Using WaitGroup
	// ------------------------------
	var wg sync.WaitGroup

	fmt.Println("Starting Goroutines with WaitGroup...")

	// Add the number of goroutines we are going to wait for
	wg.Add(2)

	// Start goroutines and pass pointer to WaitGroup
	go printNumbers(&wg)
	go printLetters(&wg)

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All Goroutines finished safely.")
}

// ------------------------------
// Key Points:
// ------------------------------
// 1. Create a WaitGroup variable: var wg sync.WaitGroup
// 2. Add number of goroutines: wg.Add(n)
// 3. Pass pointer to goroutine function: go func(&wg)
// 4. In goroutine, call wg.Done() when finished (usually with defer).
// 5. Wait for all to finish: wg.Wait()
// 6. Advantage: No arbitrary sleeps, safe synchronization.
