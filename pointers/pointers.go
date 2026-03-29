package main

import "fmt"

/*
--------------------------------
4. POINTER IN FUNCTION
--------------------------------
This function modifies original value
*/
func updateValue(x *int) {
	*x = 100 // dereference and modify
}

/*
--------------------------------
5. POINTER WITH STRUCT
--------------------------------
*/
type User struct {
	Name string
	Age  int
}

func main() {

	/*
		--------------------------------
		1. BASIC POINTER
		--------------------------------
	*/

	fmt.Println("===== 1. BASIC POINTER =====")

	a := 10

	var p *int = &a // pointer to a (stores address)

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Pointer p (stores address):", p)
	fmt.Println("Value at pointer (*p):", *p) // dereference

	/*
		--------------------------------
		2. MODIFY VALUE USING POINTER
		--------------------------------
	*/

	fmt.Println("\n===== 2. MODIFY USING POINTER =====")

	*p = 20 // changes original variable

	fmt.Println("Updated a:", a)

	/*
		--------------------------------
		3. POINTER WITHOUT 'var'
		--------------------------------
	*/

	fmt.Println("\n===== 3. SHORT SYNTAX =====")

	b := 50
	pb := &b

	fmt.Println("b:", b)
	fmt.Println("pb:", pb)
	fmt.Println("*pb:", *pb)

	/*
		--------------------------------
		4. POINTER IN FUNCTION
		--------------------------------
	*/

	fmt.Println("\n===== 4. POINTER IN FUNCTION =====")

	x := 10
	updateValue(&x) // pass address

	fmt.Println("Updated x:", x)

	/*
		--------------------------------
		5. POINTER WITH STRUCT (VERY IMPORTANT)
		--------------------------------
	*/

	fmt.Println("\n===== 5. POINTER WITH STRUCT =====")

	user := User{Name: "Jai", Age: 25}

	userPtr := &user

	userPtr.Age = 30 // no need (*userPtr).Age (Go handles it)

	fmt.Println("Updated user:", user)

	fmt.Println("\nProgram finished 🚀")
}
