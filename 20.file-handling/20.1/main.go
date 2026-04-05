package main

import (
	"fmt"
	"os"
)

// Create a file and write content to it
func main() {
	// os.Create creates or overwrites a file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // always close files

	file.WriteString("Hello, Go File Handling!\n")
	file.WriteString("Second line.\n")

	fmt.Println("File created and written successfully")
}
