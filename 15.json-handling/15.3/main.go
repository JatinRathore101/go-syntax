package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// JSON arrays: decode and encode slices of structs
func main() {
	// JSON array → slice of structs
	usersJSON := `[{"id":1,"name":"A"},{"id":2,"name":"B"}]`
	var users []User
	json.Unmarshal([]byte(usersJSON), &users)
	fmt.Println("Decoded array:", users)
	fmt.Println("First user:", users[0].Name)

	// Slice of structs → JSON array
	out, _ := json.Marshal(users)
	fmt.Println("Encoded array JSON:", string(out))
}
