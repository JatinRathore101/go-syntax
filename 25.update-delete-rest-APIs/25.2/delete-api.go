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

// DELETE /users/{id}: remove a user from the in-memory store
func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(idStr)

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"message": "user deleted"})
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/users/", deleteUser)
	fmt.Println("DELETE API running at http://localhost:8080")
	fmt.Println("Try: DELETE /users/1")
	http.ListenAndServe(":8080", nil)
}
