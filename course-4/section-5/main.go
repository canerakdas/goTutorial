package main

import "fmt"

func main() {
	// Defer is the waiting the end of main function after then run function
	defer foo()
	bar()
	// Out is: bar, foo
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
