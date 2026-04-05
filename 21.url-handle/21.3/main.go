package main

import (
	"fmt"
	"net/url"
	"path"
)

// Build a URL from scratch, join paths, and escape/unescape strings
func main() {
	// Build URL from scratch
	base := &url.URL{
		Scheme: "https",
		Host:   "api.example.com",
		Path:   "/v1/users",
	}
	q := base.Query()
	q.Set("limit", "10")
	q.Set("sort", "desc")
	base.RawQuery = q.Encode()
	fmt.Println("Built URL:", base.String())

	// Join URL path segments safely
	fullPath := path.Join("/api", "v1", "users", "123")
	fmt.Println("Joined path:", fullPath)

	// Escape special characters for use in a URL
	escaped := url.QueryEscape("jai kumar & dev")
	fmt.Println("Escaped:", escaped)

	unescaped, _ := url.QueryUnescape(escaped)
	fmt.Println("Unescaped:", unescaped)
}
