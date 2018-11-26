package main

import "fmt"

func main() {
	m := map[string]int{
		"James": 43,
		"Barba": 27,
	}

	fmt.Println(m)

	m["Todd"] = 43

	fmt.Println(m)

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v)
	}

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Caner")
	fmt.Println(m)

	if v, ok := m["Barba"]; ok {
		fmt.Println("value:", v)
		delete(m, "Barba")
		fmt.Println("Barba deleted")
	}

	fmt.Println(m)
}
