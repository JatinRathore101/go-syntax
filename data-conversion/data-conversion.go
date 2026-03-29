package main

import (
	"fmt"
	"strconv" // package for string conversions
)

func main() {

	/*
		--------------------------------
		1. BASIC TYPE CASTING
		--------------------------------
	*/

	fmt.Println("===== 1. BASIC TYPE CASTING =====")

	var a int = 10
	var b float64 = float64(a) // int → float

	fmt.Println("int:", a)
	fmt.Println("float:", b)

	/*
		--------------------------------
		2. FLOAT → INT
		--------------------------------
	*/

	fmt.Println("\n===== 2. FLOAT TO INT =====")

	var x float64 = 9.8
	var y int = int(x) // decimal part removed

	fmt.Println("float:", x)
	fmt.Println("int:", y)

	/*
		--------------------------------
		3. STRING → INT (VERY IMPORTANT)
		--------------------------------
	*/

	fmt.Println("\n===== 3. STRING TO INT =====")

	str := "123"

	num, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Converted int:", num)
	}

	/*
		--------------------------------
		4. INT → STRING
		--------------------------------
	*/

	fmt.Println("\n===== 4. INT TO STRING =====")

	n := 456
	str2 := strconv.Itoa(n)

	fmt.Println("String:", str2)

	/*
		--------------------------------
		5. STRING → FLOAT
		--------------------------------
	*/

	fmt.Println("\n===== 5. STRING TO FLOAT =====")

	strFloat := "3.14"

	f, err := strconv.ParseFloat(strFloat, 64)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Float:", f)
	}

	/*
		--------------------------------
		6. FLOAT → STRING
		--------------------------------
	*/

	fmt.Println("\n===== 6. FLOAT TO STRING =====")

	f2 := 6.28
	str3 := strconv.FormatFloat(f2, 'f', 2, 64)

	fmt.Println("String:", str3)

	fmt.Println("\nProgram finished 🚀")
}
