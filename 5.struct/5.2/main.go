package main

import "fmt"

// Accessing, updating struct fields, and using pointers to structs
type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "Jai", Age: 25}

	// Access fields
	fmt.Println("Name:", user.Name)
	fmt.Println("Age:", user.Age)

	// Update a field directly
	user.Age = 26
	fmt.Println("Updated Age:", user.Age)

	// Pointer to struct: changes affect the original
	userPtr := &user
	userPtr.Age = 30 // Go auto-dereferences: no need for (*userPtr).Age
	fmt.Println("After pointer update:", user)
}
