package main

import "fmt"

// Range loop: iterate over slices, arrays, and maps
func main() {
	nums := []int{10, 20, 30}

	// Range with index and value
	fmt.Println("Index + Value:")
	for index, value := range nums {
		fmt.Println(" Index:", index, "Value:", value)
	}

	// Range with value only (ignore index using _)
	fmt.Println("Value only:")
	for _, value := range nums {
		fmt.Println(" ", value)
	}

	// Range over a map
	scores := map[string]int{"math": 90, "science": 85}
	fmt.Println("Map range:")
	for key, val := range scores {
		fmt.Println(" ", key, "->", val)
	}
}
