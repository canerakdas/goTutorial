package main

import "fmt"

func main() {
	x := []string{"1", "2", "3"}
	y := []string{"x", "y", "z"}

	//	xy := [][]string(x, y)

	fmt.Println(x)
	fmt.Println(y)

	xy := [][]string{x, y}
	fmt.Println(xy)
}
