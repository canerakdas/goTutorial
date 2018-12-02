package main

import "fmt"

// Variatic parameters

func main() {
	foo(1, 2, 3, 4, 5, 1, 3, 4, 5)
}

func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x) // Slices of int []int

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, "-", v, "-", sum)
	}
}
