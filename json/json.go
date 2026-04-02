package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct (like a class / object schema)
// Use `json:"field_name"` to map JSON keys to struct fields
type User struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Tags  []string `json:"tags"` // array of strings
}

// Nested struct example
type Post struct {
	Title string `json:"title"`
	User  User   `json:"user"` // nested object
	Likes int    `json:"likes"`
}

// ----------- MAIN FUNCTION -----------
func main() {

	// ===============================
	// 1. JSON STRING → STRUCT (Unmarshal)
	// ===============================
	jsonStr := `{"id":1,"name":"Jai","email":"jai@test.com","tags":["dev","golang"]}`

	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Parsed Struct:", user)
	fmt.Println("User Name:", user.Name)

	// ===============================
	// 2. STRUCT → JSON (Marshal)
	// ===============================
	newUser := User{
		ID:    2,
		Name:  "Rahul",
		Email: "rahul@test.com",
		Tags:  []string{"backend", "api"},
	}

	jsonBytes, _ := json.Marshal(newUser)
	fmt.Println("JSON Output:", string(jsonBytes))

	// Pretty JSON (important for readability)
	prettyJSON, _ := json.MarshalIndent(newUser, "", "  ")
	fmt.Println("Pretty JSON:\n", string(prettyJSON))

	// ===============================
	// 3. ARRAY OF OBJECTS (Slice of structs)
	// ===============================
	users := []User{
		{ID: 1, Name: "A", Email: "a@test.com"},
		{ID: 2, Name: "B", Email: "b@test.com"},
	}

	usersJSON, _ := json.Marshal(users)
	fmt.Println("Array JSON:", string(usersJSON))

	// JSON → slice of structs
	jsonArray := `[{"id":1,"name":"A"},{"id":2,"name":"B"}]`

	var userList []User
	json.Unmarshal([]byte(jsonArray), &userList)

	fmt.Println("Parsed Array:", userList)

	// ===============================
	// 4. DYNAMIC JSON (map[string]interface{})
	// ===============================
	// Use when structure is unknown
	var data map[string]interface{}

	rawJSON := `{"name":"Jai","age":22,"isDev":true}`
	json.Unmarshal([]byte(rawJSON), &data)

	fmt.Println("Dynamic:", data)
	fmt.Println("Name:", data["name"])

	// ===============================
	// 5. ACCESS NESTED DATA
	// ===============================
	postJSON := `{
		"title": "Go Basics",
		"user": {"id": 1, "name": "Jai"},
		"likes": 10
	}`

	var post Post
	json.Unmarshal([]byte(postJSON), &post)

	fmt.Println("Post Title:", post.Title)
	fmt.Println("Author:", post.User.Name)

	// ===============================
	// 6. OPTIONAL FIELDS (omitempty)
	// ===============================
	type Product struct {
		Name  string `json:"name"`
		Price int    `json:"price,omitempty"` // omit if zero
	}

	p := Product{Name: "Laptop"}
	out, _ := json.Marshal(p)

	fmt.Println("Omit Empty:", string(out))

	// ===============================
	// 7. POINTERS (to detect null vs missing)
	// ===============================
	type Item struct {
		Name  string `json:"name"`
		Price *int   `json:"price"` // pointer allows null handling
	}

	itemJSON := `{"name":"Phone","price":null}`
	var item Item
	json.Unmarshal([]byte(itemJSON), &item)

	if item.Price == nil {
		fmt.Println("Price is NULL")
	}
}
