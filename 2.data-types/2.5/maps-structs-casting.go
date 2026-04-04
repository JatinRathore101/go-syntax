package main

import "fmt"

// Map, Struct, and Type Casting overview
type User struct {
	Name string
	Age  int
}

func main() {
	// Map: key-value pairs
	student := map[string]int{"math": 95, "science": 90}
	fmt.Println("Map:", student)
	fmt.Println("Math score:", student["math"])

	// Struct: custom object type
	user := User{Name: "Jai", Age: 25}
	fmt.Println("Struct:", user)
	fmt.Println("User Name:", user.Name)

	// Type casting: Go does NOT auto-convert types
	var intNum int = 10
	floatNum := float64(intNum) // int → float64
	fmt.Println("int to float64:", floatNum)

	var price float64 = 19.99
	whole := int(price) // float64 → int (decimal removed)
	fmt.Println("float64 to int:", whole)

	// int → string using fmt.Sprintf
	numStr := fmt.Sprintf("%d", 100)
	fmt.Println("int to string:", numStr, "| Type:", fmt.Sprintf("%T", numStr))
}
