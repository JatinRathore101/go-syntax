package main

import "fmt"

/*
--------------------------------
1. DEFINE A STRUCT
--------------------------------
Struct is like a class (but no inheritance)
*/
type User struct {
	Name string
	Age  int
}

/*
--------------------------------
5. METHOD ON STRUCT
Functions attached to struct
--------------------------------
*/
func (u User) greet() {
	fmt.Println("Hello, my name is", u.Name)
}

/*
Pointer receiver (VERY IMPORTANT)
Allows modification of original struct
*/
func (u *User) haveBirthday() {
	u.Age += 1
}

func main() {

	/*
		--------------------------------
		2. CREATE STRUCT OBJECT
		--------------------------------
	*/

	fmt.Println("===== 1. CREATE STRUCT =====")

	user1 := User{
		Name: "Jai",
		Age:  25,
	}

	fmt.Println("user1:", user1)

	/*
		--------------------------------
		3. ACCESS AND UPDATE FIELDS
		--------------------------------
	*/

	fmt.Println("\n===== 2. ACCESS / UPDATE =====")

	fmt.Println("Name:", user1.Name)

	user1.Age = 26
	fmt.Println("Updated Age:", user1.Age)

	/*
		--------------------------------
		4. POINTER TO STRUCT (VERY IMPORTANT)
		--------------------------------
	*/

	fmt.Println("\n===== 3. POINTER TO STRUCT =====")

	userPtr := &user1

	userPtr.Age = 30 // modifies original

	fmt.Println("After pointer update:", user1)

	/*
		--------------------------------
		5. METHODS
		--------------------------------
	*/

	fmt.Println("\n===== 4. METHOD CALL =====")

	user1.greet()

	/*
		--------------------------------
		6. POINTER METHOD (MODIFIES DATA)
		--------------------------------
	*/

	fmt.Println("\n===== 5. POINTER METHOD =====")

	user1.haveBirthday()

	fmt.Println("After birthday:", user1)

	fmt.Println("\nProgram finished 🚀")
}
