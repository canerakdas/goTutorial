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

type human interface {
	speak()
}

type hotdog int

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I know you are person, ", h.(person).name)
	case agent:
		fmt.Println("I know you are the agent", h.(agent).name)
	}
	fmt.Println("Hello human", h)
}

func (a agent) speak() {
	fmt.Println("Hello I'm", a.name, a.surname)
}

func (p person) speak() {
	fmt.Println("I'm a stupid person")
}

func main() {
	// p1 type of agent, but value can be more than one type (agent and human, person and human ..etc)
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

	p3 := person{
		name:    "Dr.",
		surname: "Yes",
		age:     46,
	}

	p1.speak()
	p2.speak()
	p3.speak()

	bar(p1)
	bar(p2)
	bar(p3)

	// bar able to agents and persons. poly(multiple) morph(change) ism

	// Conversion
	var x hotdog = 42

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	var y int
	y = int(x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
