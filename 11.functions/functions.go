package main

import (
	"errors"
	"fmt"
)

/*
1. SIMPLE FUNCTION (NO PARAMETERS)
*/
func greet() {
	fmt.Println("Hello! Welcome to Go functions.")
}

/*
2. FUNCTION WITH PARAMETERS
*/
func greetUser(name string) {
	fmt.Println("Hello,", name)
}

/*
3. FUNCTION WITH RETURN VALUE
*/
func add(a int, b int) int {
	return a + b
}

/*
4. FUNCTION WITH MULTIPLE RETURN VALUES
Very common in Go
*/
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

/*
5. FUNCTION RETURNING VALUE + ERROR
This pattern is EXTREMELY common in Go
*/
func safeDivide(a int, b int) (int, error) { // ********************************************************************

	if b == 0 {
		fmt.Println("errors is a package")
		return 0, errors.New("cannot divide by zero") // HERE ERRORS IS A PACKAGE
	}

	return a / b, nil
}

/*
6. NAMED RETURN VALUES
Return variables defined in signature
*/
func rectangleStats(length int, width int) (area int, perimeter int) {
	// *********************************************************************************************************************
	area = length * width
	perimeter = 2 * (length + width)

	return
}

/*
7. VARIADIC FUNCTION
Accepts unlimited arguments
*/
func sum(numbers ...int) int {

	total := 0

	for _, num := range numbers {
		total += num
	}

	return total
}

/*
8. PASS BY VALUE
Original variable will NOT change
*/
func changeValue(x int) {
	x = 100
	fmt.Println("Inside function (value):", x)
}

/*
9. PASS BY POINTER
Original variable WILL change
*/
func changeValuePointer(x *int) {
	*x = 100
}

func main() {

	fmt.Println("===== BASIC FUNCTION =====")
	greet()

	fmt.Println("\n===== FUNCTION WITH PARAMETER =====")
	greetUser("Jai")

	fmt.Println("\n===== FUNCTION WITH RETURN VALUE =====")
	result := add(5, 3)
	fmt.Println("5 + 3 =", result)

	fmt.Println("\n===== MULTIPLE RETURN VALUES =====")
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q)
	fmt.Println("Remainder:", r)

	fmt.Println("\n===== RETURN WITH ERROR =====")

	value, err := safeDivide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", value)
	}

	/*
		Try dividing by zero
	*/
	_, err2 := safeDivide(10, 0)

	if err2 != nil {
		fmt.Println("Error:", err2)
	}

	fmt.Println("\n===== NAMED RETURN VALUES =====")

	area, perimeter := rectangleStats(10, 5)
	fmt.Println("Area:", area)
	fmt.Println("Perimeter:", perimeter)

	fmt.Println("\n===== VARIADIC FUNCTION =====")

	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum:", total)

	fmt.Println("\n===== PASS BY VALUE =====")

	num := 10
	fmt.Println("Before function:", num)

	changeValue(num)

	fmt.Println("After function:", num)

	fmt.Println("\n===== PASS BY POINTER =====")

	changeValuePointer(&num)

	fmt.Println("After pointer function:", num)

	fmt.Println("\n===== ANONYMOUS FUNCTION =====")

	func() {
		fmt.Println("This is an anonymous function")
	}()

	fmt.Println("\n===== FUNCTION AS VARIABLE =====")

	multiply := func(a int, b int) int {
		return a * b
	}

	fmt.Println("Multiply 4 * 5 =", multiply(4, 5))

	fmt.Println("\nProgram finished 🚀")
}
