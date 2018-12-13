package main

import (
	"fmt"
	"net/http"
)

// Handler:
// type Handler interface{
//   ServeHTTP(ResponseWriter, *Request)
// }

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello im func")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
