package main

import "fmt"

// Methods on structs: value receiver and pointer receiver
type User struct {
	Name string
	Age  int
}

// Value receiver: reads data, does NOT modify original
func (u User) greet() {
	fmt.Println("Hello, my name is", u.Name)
}

// Pointer receiver: modifies the original struct
func (u *User) haveBirthday() {
	u.Age += 1
}

func main() {
	user := User{Name: "Jai", Age: 25}

	// Call value receiver method
	user.greet()

	// Call pointer receiver method (modifies user.Age)
	user.haveBirthday()
	fmt.Println("After birthday:", user)
}
