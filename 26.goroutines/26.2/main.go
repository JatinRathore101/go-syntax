package main

import (
	"fmt"
	"time"
)

// Multiple goroutines running concurrently
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number:", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLetters() {
	for _, l := range "ABCDE" {
		fmt.Println("Letter:", string(l))
		time.Sleep(700 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Starting two goroutines concurrently...")

	go printNumbers() // runs concurrently
	go printLetters() // runs concurrently

	// Both goroutines interleave their output
	time.Sleep(5 * time.Second)

	fmt.Println("Both goroutines finished.")
}
