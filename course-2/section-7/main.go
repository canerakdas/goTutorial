package main

import "fmt"

func main() {
	// Reduce the memory, slice arrays with make, (array,size, append size)
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3)
	x = append(x, 3)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// cap size 12*2 = 24
	// If you know the max size the slice arrays, use make :)
}
