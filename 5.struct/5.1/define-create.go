package main

import "fmt"

// Define and create a struct (like a class without inheritance)
type User struct {
	Name string
	Age  int
}

func main() {
	// Create a struct instance using field names
	user1 := User{
		Name: "Jai",
		Age:  25,
	}
	fmt.Println("user1:", user1)
	fmt.Println("Name:", user1.Name)
	fmt.Println("Age:", user1.Age)

	// Create with positional values (order must match struct definition)
	user2 := User{"Rahul", 30}
	fmt.Println("user2:", user2)
}
