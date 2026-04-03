package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ---- MODEL ----
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ---- HANDLER (POST /users) ----
func createUser(w http.ResponseWriter, r *http.Request) {

	// allow only POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// decode JSON request body → struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// (normally you would save to DB here)

	// prepare response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201

	// send JSON response
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "user created",
		"data":    user,
	})
}

// ---- MAIN ----
func main() {

	// register route
	http.HandleFunc("/users", createUser)

	fmt.Println("Server running at http://localhost:8080")

	// start server
	http.ListenAndServe(":8080", nil)
}
