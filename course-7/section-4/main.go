package main

import (
	"fmt"
	"net/http"
)

type handler int

func (m handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Caner", "Hey bla bla")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Println("Hello from the other side")
}

func main() {
	var h handler
	http.ListenAndServe(":8080", h)
}
