package main

import (
	"io"
	"net/http"
)

func dogPath(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func catPath(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func defaultPath(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "default default default")
}

func main() {
	http.HandleFunc("/dog", dogPath)
	http.HandleFunc("/cat", catPath)
	http.HandleFunc("/", defaultPath)
	http.ListenAndServe(":8080", nil)
}
