package main

import "fmt"

func main() {
	func(x int) {
		fmt.Println(x)
	}(42)

	x := func(x int) {
		fmt.Println(x)
	}

	x(33)
}
