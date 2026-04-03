package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// -----------------------------------------
	// 1. CREATE / WRITE TO A FILE
	// -----------------------------------------
	// os.Create creates (or overwrites) a file

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // always close files

	// Write string to file
	file.WriteString("Hello, Go File Handling!\n")
	file.WriteString("Second line.\n")

	fmt.Println("File created and written successfully")

	// -----------------------------------------
	// 2. APPEND TO A FILE
	// -----------------------------------------
	// Open file in append mode

	file2, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file2.Close()

	file2.WriteString("Appended line.\n")

	// -----------------------------------------
	// 3. READ FULL FILE (simple way)
	// -----------------------------------------
	// os.ReadFile reads entire file into memory

	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("\nFull file content:\n", string(data))

	// -----------------------------------------
	// 4. READ FILE LINE BY LINE (common pattern)
	// -----------------------------------------
	file3, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file3.Close()

	scanner := bufio.NewScanner(file3)

	fmt.Println("\nReading line by line:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading lines:", err)
	}

	// -----------------------------------------
	// 5. CHECK IF FILE EXISTS
	// -----------------------------------------
	_, err = os.Stat("example.txt")
	if err == nil {
		fmt.Println("\nFile exists")
	} else if os.IsNotExist(err) {
		fmt.Println("\nFile does NOT exist")
	}

	// -----------------------------------------
	// 6. DELETE A FILE
	// -----------------------------------------
	err = os.Remove("temp.txt") // deleting dummy file
	if err != nil {
		fmt.Println("temp.txt not found (ok)")
	} else {
		fmt.Println("temp.txt deleted")
	}
}
