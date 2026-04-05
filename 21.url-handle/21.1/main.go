package main

import (
	"fmt"
	"net/url"
)

// Parse a URL and read its components
func main() {
	rawURL := "https://example.com/users?id=123&name=jai"

	parsed, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsed.Scheme) // https
	fmt.Println("Host:", parsed.Host)     // example.com
	fmt.Println("Path:", parsed.Path)     // /users

	// Read query parameters
	params := parsed.Query()
	fmt.Println("id:", params.Get("id"))
	fmt.Println("name:", params.Get("name"))
}
