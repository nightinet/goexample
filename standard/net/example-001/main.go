package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func listen(addr string) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	recv, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recv from client: %s", recv)

	fmt.Fprint(conn, "Hello, Client!\n")
	if err != nil {
		log.Fatal(err)
	}
}

func connect(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprint(conn, "Hello, Server!\n")

	recv, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recv from server: %s", recv)
}

func main() {
	addr := net.JoinHostPort("", "8080")

	go listen(addr)

	time.Sleep(1 * time.Millisecond)
	connect(addr)
}
