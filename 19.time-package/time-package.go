package main

import (
	"fmt"
	"time" // important package
)

func main() {

	/*
		--------------------------------
		1. CURRENT TIME
		--------------------------------
	*/

	fmt.Println("===== 1. CURRENT TIME =====")

	now := time.Now()

	fmt.Println("Now:", now)

	/*
		--------------------------------
		2. FORMAT TIME (VERY IMPORTANT)
		Go uses a special reference date:
		2006-01-02 15:04:05
		--------------------------------
	*/

	fmt.Println("\n===== 2. FORMAT TIME =====")

	formatted := now.Format("2006-01-02 15:04:05")

	fmt.Println("Formatted:", formatted)

	/*
		--------------------------------
		3. PARSE STRING → TIME
		--------------------------------
	*/

	fmt.Println("\n===== 3. PARSE STRING =====")

	timeStr := "2025-12-25 10:30:00"

	t, err := time.Parse("2006-01-02 15:04:05", timeStr)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed time:", t)
	}

	/*
		--------------------------------
		4. ADD / SUBTRACT TIME
		--------------------------------
	*/

	fmt.Println("\n===== 4. ADD / SUBTRACT =====")

	future := now.Add(2 * time.Hour)
	past := now.Add(-24 * time.Hour)

	fmt.Println("2 hours later:", future)
	fmt.Println("24 hours ago:", past)

	/*
		--------------------------------
		5. DIFFERENCE BETWEEN TIMES
		--------------------------------
	*/

	fmt.Println("\n===== 5. TIME DIFFERENCE =====")

	diff := future.Sub(now)

	fmt.Println("Difference:", diff) // duration

	/*
		--------------------------------
		6. SLEEP (DELAY)
		--------------------------------
	*/

	fmt.Println("\n===== 6. SLEEP =====")

	fmt.Println("Sleeping for 2 seconds...")
	time.Sleep(2 * time.Second)

	fmt.Println("Awake!")

	fmt.Println("\nProgram finished 🚀")
}
