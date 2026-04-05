package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// POST /users: decode JSON body, respond with 201 Created
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "user created",
		"data":    user,
	})
}

func main() {
	http.HandleFunc("/users", createUser)
	fmt.Println("POST API running at http://localhost:8080")
	fmt.Println("Try: POST /users with JSON body {\"name\":\"Jai\",\"email\":\"jai@example.com\"}")
	http.ListenAndServe(":8080", nil)
}
