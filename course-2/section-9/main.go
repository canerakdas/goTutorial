package main

import "fmt"

// Maps

func main() {
	m := map[string]int{
		"James": 32,
		"Penny": 27,
	}
	// variable := map[Key type]Value type{}
	fmt.Println(m)
	fmt.Println(m["James"])

	fmt.Println(m["Caner"])
	// Return 0, because of default int value is 0

	v, ok := m["Caner"]
	// value, return value
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Caner"]; ok {
		fmt.Println("This is the if print", v)
	}
}
