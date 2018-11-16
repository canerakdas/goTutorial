package main

import "fmt"

var y = 42
var z = "Hello"
var p = `Hello "Wo"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	//z = 23
	// You can't assign the int value for the string declared variables, because of go static programing language.
}
