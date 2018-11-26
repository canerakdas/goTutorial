package main

import "fmt"

func main() {

	jb := []string{"James", "Bond", "Chlcolate", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Penny", "Strawberry", "martini"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
