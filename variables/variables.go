package main

import "fmt"

/*
GLOBAL VARIABLES
These are accessible from anywhere in this file
*/

var globalVar string = "I am a global variable"

// constants cannot be changed after declaration
const pi float64 = 3.14159

// globalVar2 := "gloabl variable 2" // gives syntax error here outside // syntax error: non-declaration statement outside function body

func main() {

	fmt.Println("------ GLOBAL VARIABLE ------")
	fmt.Println(globalVar)

	fmt.Println("\n------ CONSTANT ------")
	fmt.Println("Value of PI:", pi)

	/*
		----------------------------------------------------
		1. USING var (EXPLICIT TYPE)
		----------------------------------------------------
	*/

	var age int = 25
	fmt.Println("\nExplicit type variable age:", age)

	/*
		----------------------------------------------------
		2. USING var (TYPE INFERENCE)
		Go automatically detects the type
		----------------------------------------------------
	*/

	var name = "Jai"
	fmt.Println("Type inferred variable name:", name)

	/*
		----------------------------------------------------
		3. SHORT DECLARATION := (MOST COMMON IN GO)
		Works only inside functions // *********************************************
		----------------------------------------------------
	*/

	city := "Bangalore"
	fmt.Println("Short declaration city:", city, "************** Works only inside functions")

	/*
		NOTE:
		Go DOES NOT HAVE "let" keyword like JavaScript.
		Only var and :=
	*/

	/*
		----------------------------------------------------
		4. DECLARING MULTIPLE VARIABLES
		----------------------------------------------------
	*/

	var a, b int = 10, 20
	fmt.Println("\nMultiple variables:", a, b)

	/*
		You can also do this
	*/

	x, y := 100, 200
	fmt.Println("Multiple short declaration:", x, y)

	/*
		----------------------------------------------------
		5. DECLARING MULTIPLE VARIABLES (BLOCK STYLE)
		----------------------------------------------------
	*/

	var (
		username string = "admin"
		password string = "secret"
	)

	fmt.Println("\nBlock variable username:", username)
	fmt.Println("Block variable password:", password)

	/*
		----------------------------------------------------
		6. ZERO VALUES *******************************************************
		If you declare but don't initialize,
		Go gives default values.
		----------------------------------------------------
	*/

	var num int
	var text string
	var flag bool

	fmt.Println("\nZero value int:", num)
	fmt.Println("Zero value string:", text)
	fmt.Println("Zero value bool:", flag)

	/*
		----------------------------------------------------
		7. VARIABLE SCOPE
		Variables exist only within their block.
		----------------------------------------------------
	*/

	scopeVar := "I am in main()"
	fmt.Println("\nMain scope variable:", scopeVar)

	{
		blockVar := "I exist only in this block"
		fmt.Println("Block scope variable:", blockVar)
	}

	// Uncommenting below line will cause compilation error
	// fmt.Println(blockVar)

	/*
		----------------------------------------------------
		8. VARIABLE SHADOWING *********************************************
		A variable inside a block can override
		a variable from outer scope.
		----------------------------------------------------
	*/

	value := 10
	fmt.Println("\nOuter value:", value)

	{
		value := 20 // shadows outer variable
		fmt.Println("Inner value:", value)
	}

	fmt.Println("Outer value again:", value)

	/*
		----------------------------------------------------
		9. REASSIGNING VARIABLES ***********************************************
		----------------------------------------------------
	*/

	score := 50
	fmt.Println("\nInitial score:", score)

	score = 75
	fmt.Println("Updated score:", score)

	/*
		----------------------------------------------------
		10. TYPE SAFETY
		Go does NOT allow mixing types automatically
		----------------------------------------------------
	*/

	var intNum int = 10
	var floatNum float64 = 3.14

	fmt.Println("\nInteger:", intNum)
	fmt.Println("Float:", floatNum)

	// This would cause error:
	// intNum = floatNum

	fmt.Println("\nProgram finished successfully 🚀")
}
