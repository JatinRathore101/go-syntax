package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// PUT /users/{id}: update name and email of an existing user
func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	var updated User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	for i, u := range users {
		if u.ID == id {
			users[i].Name = updated.Name
			users[i].Email = updated.Email
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/users/", updateUser)
	fmt.Println("PUT API running at http://localhost:8080")
	fmt.Println("Try: PUT /users/1 with JSON body {\"name\":\"Updated\",\"email\":\"new@example.com\"}")
	http.ListenAndServe(":8080", nil)
}
