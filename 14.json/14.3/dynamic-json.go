package main

import (
	"encoding/json"
	"fmt"
)

// Dynamic JSON: use map[string]interface{} when structure is unknown
func main() {
	raw := `{"name":"Jai","age":22,"isDev":true}`

	var data map[string]interface{}
	json.Unmarshal([]byte(raw), &data)

	fmt.Println("Dynamic map:", data)
	fmt.Println("Name:", data["name"])
	fmt.Println("Age:", data["age"])
	fmt.Println("isDev:", data["isDev"])
}
