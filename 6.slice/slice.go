package main

import (
	"fmt"
)

func main() {

	fmt.Println("===== 1. CREATING A SLICE =====")

	// Most common way
	numbers := []int{10, 20, 30}

	fmt.Println("numbers:", numbers)
	fmt.Println("length:", len(numbers))
	fmt.Println("capacity:", cap(numbers))

	/*
		--------------------------------
		2. ACCESSING ELEMENTS
		--------------------------------
	*/

	fmt.Println("\n===== 2. ACCESSING ELEMENTS =====")

	fmt.Println("First element:", numbers[0])
	fmt.Println("Second element:", numbers[1])

	/*
		--------------------------------
		3. APPEND (MOST IMPORTANT)
		Add elements to slice
		--------------------------------
	*/

	fmt.Println("\n===== 3. APPEND ELEMENT =====")

	numbers = append(numbers, 40)

	fmt.Println("After append:", numbers)

	/*
		Add multiple elements
	*/

	numbers = append(numbers, 50, 60)

	fmt.Println("After appending multiple:", numbers)

	/*
		--------------------------------
		4. JOINING TWO SLICES // *************************************************************************************
		--------------------------------
	*/

	fmt.Println("\n===== 4. JOINING SLICES =====")

	moreNumbers := []int{70, 80, 90}

	// ... expands slice
	numbers = append(numbers, moreNumbers...) // ********************** // SPREAD WORKS OPPOSITE OF THE WAY IN JS

	fmt.Println("After joining:", numbers)

	/*
		--------------------------------
		5. SLICING (SUB-SLICE)
		--------------------------------
	*/

	fmt.Println("\n===== 5. SLICING =====")

	firstThree := numbers[0:3]

	fmt.Println("First three elements:", firstThree)

	/*
		--------------------------------
		6. ITERATING OVER SLICE // *****************************************************************************************
		--------------------------------
	*/

	fmt.Println("\n===== 6. LOOPING USING range =====")

	for index, value := range numbers {
		fmt.Println("Index:", index, "Value:", value)
	}

	/*
		--------------------------------
		7. POP (REMOVE LAST ELEMENT)
		Go doesn't have built-in pop
		We use slicing
		--------------------------------
	*/

	fmt.Println("\n===== 7. POP LAST ELEMENT =====")

	fmt.Println("Before pop:", numbers)

	numbers = numbers[:len(numbers)-1]

	fmt.Println("After pop:", numbers)

	/*
		--------------------------------
		8. REMOVE ELEMENT BY INDEX // ***********************************************************************************************
		--------------------------------
	*/

	fmt.Println("\n===== 8. REMOVE ELEMENT BY INDEX =====")

	indexToRemove := 2

	numbers = append(numbers[:indexToRemove], numbers[indexToRemove+1:]...)

	fmt.Println("After removing index 2:", numbers)

	/*
		--------------------------------
		9. COPYING A SLICE
		--------------------------------
	*/

	fmt.Println("\n===== 9. COPYING A SLICE =====")

	original := []int{1, 2, 3}

	// copySlice := make([]int, len(original)) // **** // MAKE FUNCTION
	// copy(copySlice, original)
	copySlice := append([]int{}, original...)

	fmt.Println("Original:", original)
	fmt.Println("Copied:", copySlice)

	/*
		--------------------------------
		10. SLICES ARE REFERENCE TYPES
		Changing one affects the other
		--------------------------------
	*/

	fmt.Println("\n===== 10. SLICE REFERENCE BEHAVIOR =====") // ********************************************************

	a := []int{100, 200, 300}

	b := a

	b[0] = 999

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	fmt.Println("\nProgram finished 🚀")
}
