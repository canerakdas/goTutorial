package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name    string
	Surname string
	Age     int
}

type items struct {
	Person []person
	Note   []int
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
func main() {
	p1 := person{
		Name:    "caner",
		Surname: "akdaş",
		Age:     26,
	}

	p2 := person{
		Name:    "bill",
		Surname: "kennedy",
		Age:     33,
	}

	persons := []person{p1, p2}

	sd := []string{"b1", "b2", "b3"}

	md := map[string]string{
		"n": "hello",
		"b": "world",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "slice.gohtml", sd)
	if err != nil {
		log.Fatalln(err)
	}

	item := items{
		Person: persons,
		Note:   []int{41, 42, 43, 43, 44},
	}
	err = tpl.ExecuteTemplate(os.Stdout, "struct.gohtml", item)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "map.gohtml", md)
	if err != nil {
		log.Fatalln(err)
	}
}
