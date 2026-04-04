package main

import (
	"fmt"
	"time"
)

// Add/subtract time, calculate difference, and sleep
func main() {
	now := time.Now()

	// Add duration to time
	future := now.Add(2 * time.Hour)
	past := now.Add(-24 * time.Hour)
	fmt.Println("2 hours later:", future.Format("2006-01-02 15:04:05"))
	fmt.Println("24 hours ago:", past.Format("2006-01-02 15:04:05"))

	// Difference between two times
	diff := future.Sub(now)
	fmt.Println("Difference:", diff) // 2h0m0s

	// Sleep: pause execution for a duration
	fmt.Println("Sleeping 1 second...")
	time.Sleep(1 * time.Second)
	fmt.Println("Awake!")
}
