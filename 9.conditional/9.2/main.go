package main

import "fmt"

// switch statement: no break needed (unlike C/Java)
func main() {
	// Basic switch
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with multiple values per case
	month := 12
	switch month {
	case 12, 1, 2:
		fmt.Println("Winter")
	case 3, 4, 5:
		fmt.Println("Spring")
	case 6, 7, 8:
		fmt.Println("Summer")
	default:
		fmt.Println("Autumn")
	}
}
