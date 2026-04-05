package main

import (
	"fmt"
	"strconv"
)

// Converting between string and float using strconv
func main() {
	// string → float64 using strconv.ParseFloat
	strFloat := "3.14"
	f, err := strconv.ParseFloat(strFloat, 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("string to float64:", f)
	}

	// float64 → string using strconv.FormatFloat
	// 'f' = decimal format, 2 = decimal places, 64 = bit size
	f2 := 6.28
	str := strconv.FormatFloat(f2, 'f', 2, 64)
	fmt.Println("float64 to string:", str, "| Type:", fmt.Sprintf("%T", str))
}
