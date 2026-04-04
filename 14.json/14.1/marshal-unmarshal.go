package main

import (
	"encoding/json"
	"fmt"
)

// JSON tags map struct fields to JSON keys
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Unmarshal: JSON string → struct
	jsonStr := `{"id":1,"name":"Jai","email":"jai@test.com"}`
	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed struct:", user)
	fmt.Println("Name:", user.Name)

	// Marshal: struct → JSON
	newUser := User{ID: 2, Name: "Rahul", Email: "rahul@test.com"}
	jsonBytes, _ := json.Marshal(newUser)
	fmt.Println("JSON:", string(jsonBytes))

	// MarshalIndent: pretty-printed JSON
	pretty, _ := json.MarshalIndent(newUser, "", "  ")
	fmt.Println("Pretty JSON:\n", string(pretty))
}
