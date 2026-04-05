package main

import (
	"fmt"
	"net/url"
)

// Add, modify, and encode query parameters
func main() {
	rawURL := "https://example.com/users?id=123"
	parsed, _ := url.Parse(rawURL)

	params := parsed.Query()

	// Add or update a parameter
	params.Set("page", "2")

	// Add multiple values for the same key
	params.Add("tag", "go")
	params.Add("tag", "api")

	// Apply changes back to the URL
	parsed.RawQuery = params.Encode()
	fmt.Println("Updated URL:", parsed.String())
}
