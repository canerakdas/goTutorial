package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type person struct {
	name    string
	surname string
	age     int
}

func firstThree(s string) (r string) {
	r = strings.TrimSpace(s)[:3]
	return
}

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"lc": strings.ToLower,
	"ft": firstThree,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.gohtml"))
	// New => create new template
	// Func => append the functions
	// ParseGlob => parse the html file
	// This flow is important!
}

func main() {
	p1 := person{
		name:    "caner",
		surname: "akdaş",
		age:     26,
	}
	err := tpl.ExecuteTemplate(os.Stdout, "hello.gohtml", p1.name)
	if err != nil {
		log.Fatalln(err)
	}
}
