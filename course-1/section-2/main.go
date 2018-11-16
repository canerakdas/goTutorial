package main

import "fmt"

func main() {
	n, err := fmt.Println("vim-go", 26, true)
	fmt.Println(n)
	fmt.Println(err)
	// If you dont want to declare and use err variable schould write this type;
	k, _ := fmt.Println("vim-vim")
	fmt.Println(k)
	// When you try to replace _ to err, go schould be throw a error because err not using anywhere
	// package.Identifier
}
