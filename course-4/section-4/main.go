package main

import "fmt"

// Variatic parameters

func main() {
	// Final parameters couple of int
	foo("string", 1, 2, 3, 45)
}

func foo(s string, x ...int) {
	fmt.Println(s)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	fmt.Printf("%T\n", x) // Slices of int []int

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println(i, "-", v, "-", sum)
	}
}
