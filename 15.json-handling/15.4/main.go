package main

import (
	"encoding/json"
	"fmt"
)

// Dynamic JSON, omitempty, and null handling with pointers
func main() {
	// Dynamic JSON: use map when structure is unknown
	raw := `{"name":"Jai","age":22,"active":true}`
	var data map[string]interface{}
	json.Unmarshal([]byte(raw), &data)
	fmt.Println("Dynamic JSON:", data)
	fmt.Println("Name:", data["name"])

	// omitempty: skip zero-value fields
	type Product struct {
		Name  string `json:"name"`
		Price int    `json:"price,omitempty"`
	}
	p := Product{Name: "Laptop"}
	pJSON, _ := json.Marshal(p)
	fmt.Println("omitempty:", string(pJSON))

	// Pointer field: detect null vs missing
	type Item struct {
		Name  string `json:"name"`
		Price *int   `json:"price"`
	}
	itemJSON := `{"name":"Phone","price":null}`
	var item Item
	json.Unmarshal([]byte(itemJSON), &item)
	if item.Price == nil {
		fmt.Println("Price is NULL")
	}
}
