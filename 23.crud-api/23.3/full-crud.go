package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Full CRUD API: GET, POST, PUT, DELETE all in one server
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

func getUsers(w http.ResponseWriter, r *http.Request) { writeJSON(w, users) }

func getUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
	for _, u := range users {
		if u.ID == id {
			writeJSON(w, u)
			return
		}
	}
	http.NotFound(w, r)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	u.ID = len(users) + 1
	users = append(users, u)
	writeJSON(w, u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
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

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
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
	}
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Full CRUD API at http://localhost:8080")
	fmt.Println("Routes: GET/POST /users | GET/PUT/DELETE /users/{id}")
	http.ListenAndServe(":8080", nil)
}
