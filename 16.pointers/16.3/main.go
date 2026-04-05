package main

import "fmt"

// Pointer to a struct: modify struct fields via pointer
type User struct {
	Name string
	Age  int
}

func main() {
	user := User{Name: "Jai", Age: 25}

	// Get pointer to struct
	userPtr := &user

	// Go auto-dereferences: userPtr.Age is same as (*userPtr).Age
	userPtr.Age = 30
	fmt.Println("Updated user:", user)

	// Pointer to struct is commonly used in functions
	fmt.Println("Name via pointer:", userPtr.Name)
}
