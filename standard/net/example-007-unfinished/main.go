package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/sys/unix"
)

func main() {
	s, err := unix.Socket(unix.AF_INET, unix.SOCK_STREAM, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer unix.Close(s)

	sa := unix.SockaddrInet4{
		Port: 8080,
	}

	err = unix.Bind(s, &sa)
	if err != nil {
		log.Fatal(err)
	}

	f := os.NewFile(uintptr(s), "/tmp/example.sock")
	defer f.Close()

	conn, err := net.FileConn(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(conn.LocalAddr())
	len, err := fmt.Fprint(conn, "Hello")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)

	buf := make([]byte, 100)
	conn.Read(buf)
	fmt.Print(string(buf))
}
