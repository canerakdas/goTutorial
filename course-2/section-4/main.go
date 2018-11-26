package main

import "fmt"

func main() {
	fmt.Println("vim-go")
package main
  
import "fmt"

func main() {
        x := []int{4,5,7,1,22}
        fmt.Println(len(x))
        fmt.Println(x)
	fmt.Println(x[1:3])
	//5,2
	fmt.Println(x[1:])
	// 5,7,1,22
	
	for i,v := range x{
		fmt.Println(i,v)
	}

	for i := 0; i< len(x);i++{
		fmt.Println(i,x[i])
	}
}
