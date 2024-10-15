package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("cannot create listener :8080")
	}	

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(fmt.Errorf("cannot handle connection: %w", err))
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	
}