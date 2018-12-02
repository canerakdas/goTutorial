package main

import "fmt"

type person struct {
	name    string
	surname string
	age     int
}

type agent struct {
	person
	keyboard bool
}

func (a agent) speak() {
	fmt.Println("Hello I'm", a.name, a.surname)
}

func main() {
	p1 := agent{
		person: person{
			name:    "Caner",
			surname: "Akdaş",
			age:     26,
		},
		keyboard: true,
	}

	p2 := agent{
		person: person{
			name:    "Foo",
			surname: "Bar",
			age:     26,
		},
		keyboard: false,
	}

	p1.speak()
	p2.speak()
}
