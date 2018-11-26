package main

import "fmt"

// Append
func main() {
	x := []int{5, 2, 3, 44, 3}
	fmt.Println(x)
	x = append(x, 1, 2, 3)
	fmt.Println(x)

	y := []int{44, 33, 22, 11}
	x = append(x, y...)
	fmt.Println(x)
}
