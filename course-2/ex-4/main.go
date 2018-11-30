package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(x)
	x = append(x, 11)
	fmt.Println(x)
	x = append(x, 12, 13, 14)
	fmt.Println(x)
	y := []int{12, 13, 14}
	x = append(x, y...)
	fmt.Println(x)

}
