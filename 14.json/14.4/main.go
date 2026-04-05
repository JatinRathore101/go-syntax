package main

import (
	"encoding/json"
	"fmt"
)

// omitempty: skip zero-value fields in JSON output
// Pointer fields: distinguish between null and missing
func main() {
	// omitempty: field is omitted if it has zero value
	type Product struct {
		Name  string `json:"name"`
		Price int    `json:"price,omitempty"` // omitted if Price == 0
	}
	p := Product{Name: "Laptop"} // Price is 0
	out, _ := json.Marshal(p)
	fmt.Println("omitempty:", string(out)) // price not in output

	// Pointer field: can distinguish null from missing
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
}
