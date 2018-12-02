package main

import "fmt"

// Variatic parameters

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 3, 2} //Slices of int
	// Slices of int diffrent for couple of(...int) int, you schould use the xi... (many int)
	foo(xi...)
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
