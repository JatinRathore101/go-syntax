package main

import (
	"encoding/json"
	"fmt"
)

// Define struct (maps JSON ↔ Go object)
type User struct {
	ID    int    `json:"id"`    // maps "id" in JSON
	Name  string `json:"name"`  // maps "name"
	Email string `json:"email"` // maps "email"
}

func main() {

	// ======================================
	// 1. JSON DECODING (Parsing JSON → struct)
	// ======================================
	jsonStr := `{"id":1,"name":"Jai","email":"jai@test.com"}`

	var user User

	// Convert JSON string → struct
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Decoded Struct:", user)
	fmt.Println("User Name:", user.Name)

	// ======================================
	// 2. JSON ENCODING (struct → JSON)
	// ======================================
	newUser := User{
		ID:    2,
		Name:  "Rahul",
		Email: "rahul@test.com",
	}

	// Convert struct → JSON (byte array)
	jsonBytes, err := json.Marshal(newUser)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("Encoded JSON:", string(jsonBytes))

	// Pretty (formatted JSON) — useful for logs/debug
	prettyJSON, _ := json.MarshalIndent(newUser, "", "  ")
	fmt.Println("Pretty JSON:\n", string(prettyJSON))

	// ======================================
	// 3. ARRAY JSON (slice of structs)
	// ======================================
	usersJSON := `[{"id":1,"name":"A"},{"id":2,"name":"B"}]`

	var users []User

	// JSON array → slice
	json.Unmarshal([]byte(usersJSON), &users)

	fmt.Println("Decoded Array:", users)
	fmt.Println("First User:", users[0].Name)

	// Encode slice → JSON
	out, _ := json.Marshal(users)
	fmt.Println("Encoded Array JSON:", string(out))

	// ======================================
	// 4. DYNAMIC JSON (unknown structure)
	// ======================================
	raw := `{"name":"Jai","age":22,"active":true}`

	var data map[string]interface{}

	json.Unmarshal([]byte(raw), &data)

	fmt.Println("Dynamic JSON:", data)
	fmt.Println("Name:", data["name"])

	// ======================================
	// 5. OPTIONAL FIELD (omitempty)
	// ======================================
	type Product struct {
		Name  string `json:"name"`
		Price int    `json:"price,omitempty"` // skipped if 0
	}

	p := Product{Name: "Laptop"}

	pJSON, _ := json.Marshal(p)
	fmt.Println("Omit Empty Field:", string(pJSON))

	// ======================================
	// 6. NULL HANDLING (using pointers)
	// ======================================
	type Item struct {
		Name  string `json:"name"`
		Price *int   `json:"price"` // pointer allows null
	}

	itemJSON := `{"name":"Phone","price":null}`

	var item Item
	json.Unmarshal([]byte(itemJSON), &item)

	if item.Price == nil {
		fmt.Println("Price is NULL")
	}

	// ======================================
	// 7. BEST PRACTICES (quick notes)
	// ======================================
	// - Always check errors from Marshal/Unmarshal
	// - Use struct instead of map when schema is known
	// - Use map[string]interface{} only for dynamic JSON
	// - Use json tags to match API field names
}
