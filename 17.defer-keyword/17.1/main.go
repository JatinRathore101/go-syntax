package main

import "fmt"

// defer: delays execution until the surrounding function returns
func main() {
	// Single defer: runs after the rest of main
	defer fmt.Println("This runs LAST (deferred)")
	fmt.Println("This runs FIRST")
}
