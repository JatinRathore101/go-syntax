package main

import (
	"fmt"
	"time"
)

// Get current time and format it
func main() {
	now := time.Now()
	fmt.Println("Current time:", now)

	// Go uses a reference date: Mon Jan 2 15:04:05 MST 2006
	// Use this exact date as the format template
	formatted := now.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted:", formatted)

	// Other format examples
	fmt.Println("Date only:", now.Format("2006-01-02"))
	fmt.Println("Time only:", now.Format("15:04:05"))
}
