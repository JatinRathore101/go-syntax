package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Example struct for JSON requests/responses
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// ---- SIMPLE GET REQUEST ----
	resp, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("GET Response:", string(body))

	// ---- GET + JSON PARSE ----
	var user User
	json.Unmarshal(body, &user)
	fmt.Println("Parsed JSON:", user.Name, user.Email)

	// ---- CUSTOM REQUEST (with headers) ----
	req, _ := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/users/1", nil)
	req.Header.Set("Authorization", "Bearer YOUR_TOKEN")

	client := &http.Client{}
	resp2, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()

	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println("GET with headers:", string(body2))

	// ---- POST REQUEST (JSON body) ----
	newUser := User{Name: "Jai", Email: "jai@example.com"}
	jsonData, _ := json.Marshal(newUser)

	resp3, err := http.Post(
		"https://jsonplaceholder.typicode.com/users",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		panic(err)
	}
	defer resp3.Body.Close()

	postBody, _ := io.ReadAll(resp3.Body)
	fmt.Println("POST Response:", string(postBody))

	// ---- POST (ADVANCED with headers) ----
	req2, _ := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/users", bytes.NewBuffer(jsonData))
	req2.Header.Set("Content-Type", "application/json")

	resp4, err := client.Do(req2)
	if err != nil {
		panic(err)
	}
	defer resp4.Body.Close()

	body4, _ := io.ReadAll(resp4.Body)
	fmt.Println("POST with headers:", string(body4))
}
