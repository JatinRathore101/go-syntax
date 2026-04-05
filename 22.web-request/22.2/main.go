package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// POST request with JSON body, simple and with custom headers
func main() {
	newUser := User{Name: "Jai", Email: "jai@example.com"}
	jsonData, _ := json.Marshal(newUser)

	// Simple POST
	resp, err := http.Post(
		"https://jsonplaceholder.typicode.com/users",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("POST Response:", string(body))

	// POST with custom headers
	req, _ := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp2, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Println("POST with headers:", string(body2))
}
