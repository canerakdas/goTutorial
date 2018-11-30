package main

import "fmt"

func main() {
	x := map[string][]string{
		"name":    []string{"caner", "mahmut", "x"},
		"surname": []string{"x", "y", "z"},
	}

	fmt.Println(x)

	fmt.Println(x["name"])
}
