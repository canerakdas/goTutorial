package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// TCP Server example, required telnet for mac
// brew install telnet
// telnet localhost 8080

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "-- Spoink")
		fmt.Fprintf(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
