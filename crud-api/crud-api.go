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

// ---- IN-MEMORY DATABASE ----
var users = []User{
	{ID: 1, Name: "Alice", Email: "alice@example.com"},
	{ID: 2, Name: "Bob", Email: "bob@example.com"},
}

// helper to write JSON response
func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// ---- HANDLERS ----

// GET /users → get all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, users)
}

// GET /users/{id} → get one user
func getUser(w http.ResponseWriter, r *http.Request) {
	// extract id from URL path
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	for _, u := range users {
		if u.ID == id {
			writeJSON(w, u)
			return
		}
	}
	http.NotFound(w, r)
}

// POST /users → create user
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	newUser.ID = len(users) + 1 // simple auto ID
	users = append(users, newUser)

	writeJSON(w, newUser)
}

// PUT /users/{id} → update user
func updateUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	var updated User
	json.NewDecoder(r.Body).Decode(&updated)

	for i, u := range users {
		if u.ID == id {
			users[i].Name = updated.Name
			users[i].Email = updated.Email
			writeJSON(w, users[i])
			return
		}
	}
	http.NotFound(w, r)
}

// DELETE /users/{id} → delete user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	for i, u := range users {
		if u.ID == id {
			// remove element from slice
			users = append(users[:i], users[i+1:]...)
			writeJSON(w, map[string]string{"message": "deleted"})
			return
		}
	}
	http.NotFound(w, r)
}

// ---- ROUTER (basic) ----
func handler(w http.ResponseWriter, r *http.Request) {
	// route based on method + path
	if r.URL.Path == "/users" {
		switch r.Method {
		case "GET":
			getUsers(w, r)
		case "POST":
			createUser(w, r)
		default:
			http.NotFound(w, r)
		}
		return
	}

	if strings.HasPrefix(r.URL.Path, "/users/") {
		switch r.Method {
		case "GET":
			getUser(w, r)
		case "PUT":
			updateUser(w, r)
		case "DELETE":
			deleteUser(w, r)
		default:
			http.NotFound(w, r)
		}
		return
	}

	http.NotFound(w, r)
}

// ---- MAIN ----
func main() {
	http.HandleFunc("/", handler)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
