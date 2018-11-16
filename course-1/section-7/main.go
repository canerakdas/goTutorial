// GO IS ABOUT TYPE!
// TYPE IS LIFE!
package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b = 43
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
