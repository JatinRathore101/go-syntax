package main

import (
	"bufio"
	"fmt"
	"os"
)

// Append to a file, read full content, and read line by line
func main() {
	// Append to existing file
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	file.WriteString("Appended line.\n")
	file.Close()

	// Read full file at once
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Full content:\n", string(data))

	// Read line by line using bufio.Scanner
	f, _ := os.Open("example.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	fmt.Println("Line by line:")
	for scanner.Scan() {
		fmt.Println(" ", scanner.Text())
	}
}
