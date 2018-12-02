package main

import "fmt"

func main() {
	x := fact(4)
	fmt.Println(x)
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
