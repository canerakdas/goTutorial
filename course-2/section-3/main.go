package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 1, 22}
	fmt.Println(len(x))
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}
}
