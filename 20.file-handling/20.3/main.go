package main

import (
	"fmt"
	"os"
)

// Check if a file exists and delete a file
func main() {
	// Check if file exists using os.Stat
	_, err := os.Stat("example.txt")
	if err == nil {
		fmt.Println("example.txt exists")
	} else if os.IsNotExist(err) {
		fmt.Println("example.txt does NOT exist")
	}

	// Delete a file
	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Could not delete:", err)
	} else {
		fmt.Println("example.txt deleted")
	}
}
