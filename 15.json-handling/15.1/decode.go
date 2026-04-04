package main

import (
	"encoding/json"
	"fmt"
)

// JSON decoding: parse a JSON string into a Go struct
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	jsonStr := `{"id":1,"name":"Jai","email":"jai@test.com"}`

	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded struct:", user)
	fmt.Println("User Name:", user.Name)
	fmt.Println("User Email:", user.Email)
}
