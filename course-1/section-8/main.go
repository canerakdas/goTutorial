package main

import "fmt"

type hotdog int

var b hotdog

var a int

func main() {
	b = 55
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	//fmt.Printf("%T\n", main)
	// Conversions: https://golang.org/ref/spec#Conversions
	a = int(b)

	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
