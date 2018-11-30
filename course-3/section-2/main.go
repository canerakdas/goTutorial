package main

import "fmt"

type person struct {
	name    string
	surname string
}

type agent struct {
	person
	shield bool
}

func main() {

	p1 := person{
		name:    "Caner",
		surname: "Akdaş",
	}

	a1 := agent{
		person: person{
			name:    "James",
			surname: "Bond,",
		},
		shield: true,
	}

	a2 := struct {
		person
		gun bool
	}{
		person: person{
			name:    "James",
			surname: "Bond,",
		},
		gun: true,
	}

	fmt.Println(p1)

	fmt.Println(a1)

	fmt.Println(a2)
}
