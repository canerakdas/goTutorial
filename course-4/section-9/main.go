package main

import "fmt"

func main() {
	x := foo()
	fmt.Printf("%T\n", x)

	fmt.Println(x())
	// Also

	fmt.Println(foo()())

}

func foo() func() int {
	return func() int {
		return 451
	}
}
