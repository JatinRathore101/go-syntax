package main

import (
	"fmt"
	"strconv"
)

// Converting between string and int using strconv
func main() {
	// string → int using strconv.Atoi
	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("string to int:", num)
	}

	// int → string using strconv.Itoa
	n := 456
	str2 := strconv.Itoa(n)
	fmt.Println("int to string:", str2, "| Type:", fmt.Sprintf("%T", str2))

	// Invalid conversion example
	_, err2 := strconv.Atoi("abc")
	if err2 != nil {
		fmt.Println("Conversion error:", err2)
	}
}
