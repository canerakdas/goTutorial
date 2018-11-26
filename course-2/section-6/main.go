package main

import "fmt"

func main() {
	x := []int{55, 33, 22, 11, 55, 55}
	x = append(x[:2], x[4:]...)
}
