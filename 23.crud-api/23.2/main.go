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

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// POST /users → create a new user
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)
	newUser.ID = len(users) + 1
	users = append(users, newUser)
	writeJSON(w, newUser)
}

// PUT /users/{id} → update an existing user
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

// DELETE /users/{id} → delete a user
func deleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			writeJSON(w, map[string]string{"message": "deleted"})
			return
		}
	}
	http.NotFound(w, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case "POST":
			createUser(w, r)
		default:
			http.NotFound(w, r)
		}
		return
	}
	if strings.HasPrefix(r.URL.Path, "/users/") {
		switch r.Method {
		case "PUT":
			updateUser(w, r)
		case "DELETE":
			deleteUser(w, r)
		default:
			http.NotFound(w, r)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("POST/PUT/DELETE API running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
