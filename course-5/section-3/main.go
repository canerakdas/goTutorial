package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	name    string
	surname string
	age     int
}

func init() {
	tpl = template.Must(template.ParseGlob("hello.gohtml"))
}
func main() {
	p1 := person{
		name:    "caner",
		surname: "akdaş",
		age:     26,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", p1)
	if err != nil {
		log.Fatalln(err)
	}
}
