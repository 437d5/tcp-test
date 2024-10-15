package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	defer conn.Close()

	name := conn.RemoteAddr().String()

	log.Printf("%+v connected\n", name)
	conn.Write([]byte("Hello, " + name + "\n\r"))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			conn.Write([]byte("Bye\n\r"))
			log.Println(name, "disconnected")
			break
		}
	}
}