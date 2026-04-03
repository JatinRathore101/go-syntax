package main

import (
	"fmt"
	"strings" // important package
)

func main() {

	/*
		--------------------------------
		1. CONTAINS / PREFIX / SUFFIX
		--------------------------------
	*/

	fmt.Println("===== 1. CONTAINS / PREFIX / SUFFIX =====")

	text := "hello golang world"

	fmt.Println("Contains 'golang':", strings.Contains(text, "golang"))
	fmt.Println("HasPrefix 'hello':", strings.HasPrefix(text, "hello"))
	fmt.Println("HasSuffix 'world':", strings.HasSuffix(text, "world"))

	/*
		--------------------------------
		2. SPLIT AND JOIN
		--------------------------------
	*/

	fmt.Println("\n===== 2. SPLIT AND JOIN =====")

	str := "apple,banana,mango"

	parts := strings.Split(str, ",")
	fmt.Println("Split:", parts)

	joined := strings.Join(parts, " | ")
	fmt.Println("Joined:", joined)

	/*
		--------------------------------
		3. REPLACE
		--------------------------------
	*/

	fmt.Println("\n===== 3. REPLACE =====")

	sentence := "I like java"

	newSentence := strings.Replace(sentence, "java", "go", 1)

	fmt.Println("Original:", sentence)
	fmt.Println("Updated:", newSentence)

	/*
		--------------------------------
		4. TO UPPER / LOWER
		--------------------------------
	*/

	fmt.Println("\n===== 4. CASE CONVERSION =====")

	name := "GoLang"

	fmt.Println("Upper:", strings.ToUpper(name))
	fmt.Println("Lower:", strings.ToLower(name))

	/*
		--------------------------------
		5. TRIM SPACES
		--------------------------------
	*/

	fmt.Println("\n===== 5. TRIM =====")

	raw := "   hello world   "

	fmt.Println("Before:", "["+raw+"]")
	fmt.Println("After:", "["+strings.TrimSpace(raw)+"]")

	/*
		--------------------------------
		6. COUNT AND INDEX
		--------------------------------
	*/

	fmt.Println("\n===== 6. COUNT AND INDEX =====")

	data := "go go go"

	fmt.Println("Count 'go':", strings.Count(data, "go"))
	fmt.Println("Index of first 'go':", strings.Index(data, "go"))

	fmt.Println("\nProgram finished 🚀")
}
