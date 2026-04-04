package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Simple GET request and GET with custom headers
func main() {
	// Simple GET
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("GET Response:", string(body))

	// Parse JSON from response
	var user User
	json.Unmarshal(body, &user)
	fmt.Println("Parsed Name:", user.Name)

	// GET with custom headers
	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users/1", nil)
	req.Header.Set("Authorization", "Bearer <your_token>")

	client := &http.Client{}
	resp2, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println("GET with headers status:", resp2.Status)
	_ = body2
}
