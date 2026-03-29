package main

import "fmt"

func main() {

	/*
		--------------------------------
		1. CREATING A MAP
		--------------------------------
	*/

	fmt.Println("===== 1. CREATE MAP =====")

	// Most common way
	scores := map[string]int{
		"math":    90,
		"science": 85,
	}

	fmt.Println("scores:", scores)

	/*
		--------------------------------
		2. ADD / UPDATE VALUE
		--------------------------------
	*/

	fmt.Println("\n===== 2. ADD / UPDATE =====")

	// Add new key
	scores["english"] = 88

	// Update existing key
	scores["math"] = 95

	fmt.Println("updated scores:", scores)

	/*
		--------------------------------
		3. ACCESS VALUE
		--------------------------------
	*/

	fmt.Println("\n===== 3. ACCESS VALUE =====")

	fmt.Println("Math score:", scores["math"])

	/*
		If key does not exist → returns zero value
		(int → 0, string → "")
	*/
	fmt.Println("History score (not present):", scores["history"])

	/*
		--------------------------------
		4. CHECK IF KEY EXISTS (VERY IMPORTANT)
		--------------------------------
	*/

	fmt.Println("\n===== 4. CHECK KEY EXISTS =====")

	value, exists := scores["science"]

	if exists {
		fmt.Println("Science exists with value:", value)
	} else {
		fmt.Println("Science not found")
	}

	/*
		--------------------------------
		5. DELETE KEY
		--------------------------------
	*/

	fmt.Println("\n===== 5. DELETE KEY =====")

	delete(scores, "english")

	fmt.Println("after delete:", scores)

	/*
		--------------------------------
		6. LOOP OVER MAP (range)
		--------------------------------
	*/

	fmt.Println("\n===== 6. LOOP MAP =====")

	for key, value := range scores {
		fmt.Println("Subject:", key, "Score:", value)
	}

	/*
		If only keys needed
	*/
	for key := range scores {
		fmt.Println("Key only:", key)
	}

	fmt.Println("\nProgram finished 🚀")
}
