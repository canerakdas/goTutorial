package main

import "fmt"

//struct
type person struct {
	name    string
	surname string
}

func main() {
	p1 := person{
		name:    "caner",
		surname: "akdas",
	}

	p2 := person{
		name:    "c",
		surname: "a",
	}

	fmt.Println(p1)

	fmt.Println(p2.name, p2.surname)

}
