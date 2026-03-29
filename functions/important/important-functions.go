package main

import (
	"errors"
	"fmt"
)

/*
1️⃣ BASIC FUNCTION
Most common function style in Go
Takes parameters and returns a value
*/
func add(a int, b int) int {
	return a + b
}

/*
2️⃣ MULTIPLE RETURN VALUES
Go allows returning multiple values
*/
func divide(a int, b int) (int, int) {

	quotient := a / b
	remainder := a % b

	return quotient, remainder
}

/*
3️⃣ VALUE + ERROR PATTERN (VERY IMPORTANT)
Most Go libraries follow this pattern
*/
func safeDivide(a int, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {

	fmt.Println("===== 1 BASIC FUNCTION =====")

	sum := add(5, 3)
	fmt.Println("5 + 3 =", sum)

	fmt.Println("\n===== 2 MULTIPLE RETURN VALUES =====")

	q, r := divide(10, 3)

	fmt.Println("Quotient:", q)
	fmt.Println("Remainder:", r)

	fmt.Println("\n===== 3 VALUE + ERROR PATTERN =====")

	result, err := safeDivide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)

	/*
		Try dividing by zero
	*/
	fmt.Println("\nTrying division by zero...")

	_, err2 := safeDivide(10, 0)

	if err2 != nil {
		fmt.Println("Handled Error:", err2)
	}

	/*
		4️⃣ ANONYMOUS FUNCTION
		Commonly used for goroutines and quick logic
	*/

	fmt.Println("\n===== 4 ANONYMOUS FUNCTION =====")

	multiply := func(a int, b int) int {
		return a * b
	}

	fmt.Println("4 * 5 =", multiply(4, 5))

	fmt.Println("\nProgram finished 🚀")
}
