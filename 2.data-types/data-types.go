package main

import (
	"fmt"
)

/*
STRUCT
Used to create custom data types (similar to objects in other languages)
*/
type User struct {
	Name string
	Age  int
}

func main() {

	fmt.Println("===== BASIC DATA TYPES =====")

	/*
		-----------------------------------
		1. INTEGER TYPES
		-----------------------------------
	*/
	var a int = 10
	var b int8 = 120
	var c int64 = 9000000000

	fmt.Println("\nIntegers:")
	fmt.Println("a =", a, "Type:", fmt.Sprintf("%T", a))
	fmt.Println("b =", b, "Type:", fmt.Sprintf("%T", b))
	fmt.Println("c =", c, "Type:", fmt.Sprintf("%T", c))

	/*
		-----------------------------------
		2. FLOAT TYPES
		-----------------------------------
	*/
	var price float32 = 19.99
	var pi float64 = 3.1415926535

	fmt.Println("\nFloats:")
	fmt.Println("price =", price, "Type:", fmt.Sprintf("%T", price))
	fmt.Println("pi =", pi, "Type:", fmt.Sprintf("%T", pi))

	/*
		-----------------------------------
		3. BOOLEAN TYPE
		-----------------------------------
	*/
	var isActive bool = true

	fmt.Println("\nBoolean:")
	fmt.Println("isActive =", isActive, "Type:", fmt.Sprintf("%T", isActive))

	/*
		-----------------------------------
		4. STRING TYPE
		-----------------------------------
	*/
	var name string = "Jai"

	fmt.Println("\nString:")
	fmt.Println("name =", name, "Type:", fmt.Sprintf("%T", name))

	/*
		-----------------------------------
		5. ARRAYS
		Fixed size collection
		-----------------------------------
	*/
	var numbers [3]int = [3]int{10, 20, 30}

	fmt.Println("\nArray:")
	fmt.Println("numbers =", numbers)
	fmt.Println("Type:", fmt.Sprintf("%T", numbers))

	/*
		-----------------------------------
		6. SLICES (MOST USED LIST TYPE)
		Dynamic size
		-----------------------------------
	*/
	scores := []int{90, 85, 70}

	fmt.Println("\nSlice:")
	fmt.Println("scores =", scores)
	fmt.Println("Type:", fmt.Sprintf("%T", scores))

	/*
		Add element to slice
	*/
	scores = append(scores, 95)

	fmt.Println("Updated slice:", scores)

	/*
		-----------------------------------
		7. MAPS (OBJECT / DICTIONARY)
		Key-value pairs
		-----------------------------------
	*/
	student := map[string]int{
		"math":    95,
		"science": 90,
	}

	fmt.Println("\nMap:")
	fmt.Println("student =", student)
	fmt.Println("Type:", fmt.Sprintf("%T", student))

	/*
		Access map value
	*/
	fmt.Println("Math score:", student["math"])

	/*
		-----------------------------------
		8. STRUCT (CUSTOM OBJECT TYPE)
		-----------------------------------
	*/
	user := User{
		Name: "Jai",
		Age:  25,
	}

	fmt.Println("\nStruct (object-like):")
	fmt.Println("user =", user)
	fmt.Println("Type:", fmt.Sprintf("%T", user))

	fmt.Println("User Name:", user.Name)
	fmt.Println("User Age:", user.Age)

	/*
		-----------------------------------
		9. TYPE CASTING / TYPE CONVERSION
		Go does NOT auto convert types
		-----------------------------------
	*/

	fmt.Println("\n===== TYPE CASTING =====")

	var intNum int = 10
	var floatNum float64

	/*
		Convert int -> float64
	*/
	floatNum = float64(intNum)

	fmt.Println("Original int:", intNum, "Type:", fmt.Sprintf("%T", intNum))
	fmt.Println("Converted float:", floatNum, "Type:", fmt.Sprintf("%T", floatNum))

	/*
		Convert float -> int
		Note: decimal part gets removed
	*/
	var price2 float64 = 19.99
	var wholeNumber int = int(price2)

	fmt.Println("\nOriginal float:", price2)
	fmt.Println("Converted to int:", wholeNumber)

	/*
		Convert int -> string *********************8
	*/
	number := 100
	numberString := fmt.Sprintf("%d", number)

	fmt.Println("\nNumber:", number)
	fmt.Println("Number as string:", numberString, "Type:", fmt.Sprintf("%T", numberString))

	fmt.Println("\nProgram completed 🚀")
}
