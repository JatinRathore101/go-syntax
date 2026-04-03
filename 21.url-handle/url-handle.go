package main

import (
	"fmt"
	"net/url"
	"path"
)

func main() {

	// ---- PARSE A URL ----
	rawURL := "https://example.com/users?id=123&name=jai"

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Println("Scheme:", parsedURL.Scheme) // https
	fmt.Println("Host:", parsedURL.Host)     // example.com
	fmt.Println("Path:", parsedURL.Path)     // /users

	// ---- READ QUERY PARAMETERS ----
	queryParams := parsedURL.Query() // returns map[string][]string

	fmt.Println("id:", queryParams.Get("id"))     // "123"
	fmt.Println("name:", queryParams.Get("name")) // "jai"

	// ---- ADD / MODIFY QUERY PARAMETERS ----
	queryParams.Set("page", "2")  // add or update
	queryParams.Add("tag", "go")  // add multiple values
	queryParams.Add("tag", "api") // ?tag=go&tag=api

	parsedURL.RawQuery = queryParams.Encode() // apply changes
	fmt.Println("Updated URL:", parsedURL.String())

	// ---- BUILD URL FROM SCRATCH ----
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "api.example.com",
		Path:   "/v1/users",
	}

	// add query params
	q := baseURL.Query()
	q.Set("limit", "10")
	q.Set("sort", "desc")
	baseURL.RawQuery = q.Encode()

	fmt.Println("Built URL:", baseURL.String())

	// ---- HANDLE PATHS ----
	// safely join URL paths
	basePath := "/api"
	fullPath := path.Join(basePath, "v1", "users", "123")
	fmt.Println("Joined Path:", fullPath) // /api/v1/users/123

	// ---- ESCAPE / UNESCAPE ----
	// useful when URL contains special characters
	escaped := url.QueryEscape("jai kumar & dev")
	fmt.Println("Escaped:", escaped)

	unescaped, _ := url.QueryUnescape(escaped)
	fmt.Println("Unescaped:", unescaped)
}
