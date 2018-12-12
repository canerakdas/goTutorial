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

func (p person) DoubleAge() int {
	return p.Age * 2
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
	// New => create new template
	// Func => append the functions
	// ParseGlob => parse the html file
	// This flow is important!
}

func main() {
	p1 := person{
		Name:    "Caner",
		Surname: "Akdaş",
		Age:     26,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
