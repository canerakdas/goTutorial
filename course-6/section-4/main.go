package main

import (
	"fmt"
	"log"
	"net"
)

// This code running with  course 6, section 2
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
