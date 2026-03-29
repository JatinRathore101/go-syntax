package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func main() {

	// -----------------------------------------
	// 1. CASE CONVERSION
	// -----------------------------------------
	s := "Hello GoLang 123"

	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("ToUpper:", strings.ToUpper(s))

	// -----------------------------------------
	// 2. CHECK CHARACTER TYPES (using unicode)
	// -----------------------------------------
	ch := 'A' // rune (character)

	fmt.Println("IsLetter:", unicode.IsLetter(ch))
	fmt.Println("IsDigit:", unicode.IsDigit(ch))
	fmt.Println("IsSpace:", unicode.IsSpace(ch))

	// Alphanumeric = letter OR digit
	fmt.Println("IsAlphanumeric:", unicode.IsLetter(ch) || unicode.IsDigit(ch))

	// -----------------------------------------
	// 3. STRING TRIMMING (VERY COMMON)
	// -----------------------------------------
	str := "   hello world   "

	fmt.Println("TrimSpace:", strings.TrimSpace(str))      // remove spaces both sides
	fmt.Println("TrimLeft:", strings.TrimLeft(str, " "))   // left side
	fmt.Println("TrimRight:", strings.TrimRight(str, " ")) // right side

	// Remove specific characters
	fmt.Println("Trim chars:", strings.Trim("!!hello!!", "!"))

	// -----------------------------------------
	// 4. STRING SEARCH / CHECK
	// -----------------------------------------
	text := "golang is awesome"

	fmt.Println("Contains 'go':", strings.Contains(text, "go"))
	fmt.Println("HasPrefix 'go':", strings.HasPrefix(text, "go"))
	fmt.Println("HasSuffix 'me':", strings.HasSuffix(text, "me"))
	fmt.Println("Index of 'is':", strings.Index(text, "is"))

	// -----------------------------------------
	// 5. SPLIT & JOIN
	// -----------------------------------------
	parts := strings.Split("a,b,c", ",")
	fmt.Println("Split:", parts)

	joined := strings.Join(parts, "-")
	fmt.Println("Join:", joined)

	// -----------------------------------------
	// 6. CUSTOM SORTING STRINGS (IMPORTANT)
	// -----------------------------------------
	words := []string{"Go", "java", "Python", "c"}

	// Case-insensitive sorting
	sort.Slice(words, func(i, j int) bool {
		return strings.ToLower(words[i]) < strings.ToLower(words[j])
	})
	fmt.Println("Case-insensitive sort:", words)

	// -----------------------------------------
	// 7. CUSTOM SORT BY LENGTH
	// -----------------------------------------
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	fmt.Println("Sort by length:", words)

	// -----------------------------------------
	// 8. LOOP OVER STRING (IMPORTANT for chars)
	// -----------------------------------------
	for _, r := range "Go123!" {
		if unicode.IsLetter(r) {
			fmt.Println(string(r), "is Letter")
		} else if unicode.IsDigit(r) {
			fmt.Println(string(r), "is Digit")
		} else {
			fmt.Println(string(r), "is Other")
		}
	}
}
