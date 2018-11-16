package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	fmt.Println("Hello world 2")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

func bar() {
	fmt.Println("I'm in bar")
}

// Control flow
// (1) Sequence
// (2) Loop
// (3) Conditional
