package main

import (
	"errors"
	"fmt"
)

// Value + error return pattern: the most common Go pattern
func safeDivide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	// Successful division
	result, err := safeDivide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Division by zero
	_, err2 := safeDivide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	}
}
