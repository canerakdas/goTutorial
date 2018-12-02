package main

import "fmt"

func main() {
	foo()
	bar("bar")
	fmt.Println(woo("woo"))
	fmt.Println(doo("doo", "doo2"))
}

func foo() {
	fmt.Println("Hello foo")
}

func bar(s string) {
	fmt.Println("Hello ", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello ", s)
}

func doo(s string, f string) (string, bool) {
	return fmt.Sprint("Hello ", s, " ", f), false
}
