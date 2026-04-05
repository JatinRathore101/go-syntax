package main

import (
	"encoding/json"
	"fmt"
)

// JSON encoding: convert a Go struct to JSON
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	user := User{ID: 2, Name: "Rahul", Email: "rahul@test.com"}

	// Marshal: struct → compact JSON
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("Encoded JSON:", string(jsonBytes))

	// MarshalIndent: struct → pretty JSON
	pretty, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println("Pretty JSON:\n", string(pretty))
}
