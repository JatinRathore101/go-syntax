package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ---- MODEL ----
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ---- IN-MEMORY DATA ----
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// ---- UPDATE USER (PUT /users/{id}) ----
func updateUser(w http.ResponseWriter, r *http.Request) {

	// allow only PUT
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// extract id from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	// decode request body
	var updated User
	err := json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// find and update user
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

// ---- DELETE USER (DELETE /users/{id}) ----
func deleteUser(w http.ResponseWriter, r *http.Request) {

	// allow only DELETE
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// extract id from URL
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	// find and delete user
	for i, u := range users {
		if u.ID == id {
			// remove from slice
			users = append(users[:i], users[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "user deleted",
			})
			return
		}
	}

	http.NotFound(w, r)
}

// ---- ROUTER ----
func handler(w http.ResponseWriter, r *http.Request) {

	if strings.HasPrefix(r.URL.Path, "/users/") {

		if r.Method == http.MethodPut {
			updateUser(w, r)
			return
		}

		if r.Method == http.MethodDelete {
			deleteUser(w, r)
			return
		}
	}

	http.NotFound(w, r)
}

// ---- MAIN ----
func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
