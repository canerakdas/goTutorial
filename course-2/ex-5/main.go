package main

import "fmt"

func main() {
	city := make([]string, 10, 50)
	fmt.Println(city)
	fmt.Println(len(city))
	fmt.Println(cap(city))
}
