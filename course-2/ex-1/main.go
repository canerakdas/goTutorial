package main

import "fmt"

func main() {
	a := [5]int{}

	for i := range a {
		a[i] = i
		fmt.Println(a)
	}
}
