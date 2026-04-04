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

// GET /users → return all users
func getUsers(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, users)
}

// GET /users/{id} → return one user by ID
func getUser(w http.ResponseWriter, r *http.Request) {
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

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" && r.Method == "GET" {
		getUsers(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/users/") && r.Method == "GET" {
		getUser(w, r)
		return
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("GET API running at http://localhost:8080")
	fmt.Println("Try: GET /users  or  GET /users/1")
	http.ListenAndServe(":8080", nil)
}
