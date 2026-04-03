package main

import "fmt"

func main() {

	fmt.Println("===== 1. BASIC FOR LOOP =====")

	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	/*
		--------------------------------
		2. WHILE-STYLE LOOP
		Go does not have 'while', use for
		--------------------------------
	*/

	fmt.Println("\n===== 2. WHILE STYLE =====")

	j := 0

	for j < 5 {
		fmt.Println("j =", j)
		j++
	}

	/*
		--------------------------------
		3. INFINITE LOOP
		--------------------------------
	*/

	fmt.Println("\n===== 3. INFINITE LOOP (break used) =====")

	k := 0

	for {
		fmt.Println("k =", k)
		k++

		if k == 3 {
			break // exit loop
		}
	}

	/*
		--------------------------------
		4. RANGE LOOP (MOST IMPORTANT)
		Used for slices, arrays, maps
		--------------------------------
	*/

	fmt.Println("\n===== 4. RANGE LOOP =====")

	nums := []int{10, 20, 30}

	for index, value := range nums {
		fmt.Println("Index:", index, "Value:", value)
	}

	/*
		If you only need value
	*/

	for _, value := range nums {
		fmt.Println("Value only:", value)
	}

	/*
		--------------------------------
		5. BREAK AND CONTINUE
		--------------------------------
	*/

	fmt.Println("\n===== 5. CONTINUE =====")

	for i := 0; i < 5; i++ {

		if i == 2 {
			continue // skip this iteration
		}

		fmt.Println("i =", i)
	}

	fmt.Println("\n===== 6. BREAK =====")

	for i := 0; i < 5; i++ {

		if i == 3 {
			break // stop loop
		}

		fmt.Println("i =", i)
	}

	fmt.Println("\nProgram finished 🚀")
}
