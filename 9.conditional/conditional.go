package main

import "fmt"

func main() {

	fmt.Println("===== 1. SIMPLE IF =====")

	age := 20

	if age >= 18 {
		fmt.Println("You are an adult")
	}

	/*
		--------------------------------
		2. IF - ELSE
		--------------------------------
	*/

	fmt.Println("\n===== 2. IF ELSE =====")

	score := 40

	if score >= 50 {
		fmt.Println("Passed")
	} else {
		fmt.Println("Failed")
	}

	/*
		--------------------------------
		3. IF - ELSE IF LADDER
		(Go doesn't have 'elif', we use 'else if')
		--------------------------------
	*/

	fmt.Println("\n===== 3. IF ELSE IF LADDER =====")

	marks := 82

	if marks >= 90 {
		fmt.Println("Grade: A")
	} else if marks >= 75 {
		fmt.Println("Grade: B")
	} else if marks >= 50 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: Fail")
	}

	/*
		--------------------------------
		4. SWITCH STATEMENT
		Common alternative to long if-else
		--------------------------------
	*/

	fmt.Println("\n===== 4. SWITCH =====") // NO break; // needed in golang switch // instead opposite needed to fall  // ***************************8

	day := 3

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	/*
		--------------------------------
		5. SWITCH WITH MULTIPLE VALUES
		--------------------------------
	*/

	fmt.Println("\n===== 5. SWITCH MULTIPLE CASES =====")

	month := 12

	switch month {
	case 12, 1, 2:
		fmt.Println("Winter")
	case 3, 4, 5:
		fmt.Println("Spring")
	case 6, 7, 8:
		fmt.Println("Summer")
	default:
		fmt.Println("Autumn")
	}

	/*
		--------------------------------
		6. SWITCH WITHOUT CONDITION
		Common Go pattern
		Works like cleaner if-else ladder
		--------------------------------
	*/

	fmt.Println("\n===== 6. SWITCH WITHOUT CONDITION =====")

	temp := 28

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

	fmt.Println("\nProgram finished 🚀")
}
