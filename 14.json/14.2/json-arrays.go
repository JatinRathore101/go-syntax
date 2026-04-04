package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// JSON arrays: marshal and unmarshal slices of structs
func main() {
	// JSON array → slice of structs
	jsonArray := `[{"id":1,"name":"A","email":"a@test.com"},{"id":2,"name":"B","email":"b@test.com"}]`
	var users []User
	json.Unmarshal([]byte(jsonArray), &users)
	fmt.Println("Parsed array:", users)
	fmt.Println("First user:", users[0].Name)

	// Slice of structs → JSON array
	out, _ := json.Marshal(users)
	fmt.Println("Encoded array:", string(out))

	// Nested struct
	type Post struct {
		Title string `json:"title"`
		User  User   `json:"user"`
		Likes int    `json:"likes"`
	}
	postJSON := `{"title":"Go Basics","user":{"id":1,"name":"Jai"},"likes":10}`
	var post Post
	json.Unmarshal([]byte(postJSON), &post)
	fmt.Println("Post title:", post.Title)
	fmt.Println("Author:", post.User.Name)
}
