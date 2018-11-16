package main

import "fmt"

// non-declaration statement outside function body
// if you try to declare variable here with short declaration operator; x := 5

// Declare & assign
var z = 3

// Declare variable and assign the zero value
// false for booleans, 0 for numeric types, "" for strings, and nil for pointers, functions, interfaces, slices, channels, and maps.
var t int
var s string
var b bool

func main() {
	// Short decleration operator :=
	// Declare a variable and assign a value (of a certain type)
	x := 42
	fmt.Println(x)

	var y = 43
	fmt.Println(y)

	fmt.Println(z)

	foo()

	fmt.Println(t)
}

func foo() {
	fmt.Println("again:", z)
}

// Short decleration should be using function scope.
// Var keyword decleration also usable the global scope.
