package main

import (
	"fmt"
	"time"
)

// Parse a string into a time.Time value
func main() {
	timeStr := "2025-12-25 10:30:00"

	// Layout must match the reference date format
	t, err := time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Parsed time:", t)
	fmt.Println("Year:", t.Year())
	fmt.Println("Month:", t.Month())
	fmt.Println("Day:", t.Day())
}
