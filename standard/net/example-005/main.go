package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	p1, p2 := net.Pipe()

	go func() {
		wbuf := []byte("Hello!")
		_, err := p1.Write(wbuf)
		if err != nil {
			log.Fatalln(err)
		}

		rbuf := make([]byte, 1024)
		_, err = p2.Read(rbuf)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(string(rbuf))
	}()

	rbuf := make([]byte, 1024)
	_, err := p2.Read(rbuf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rbuf))

	wbuf := []byte("Hello again!")
	_, err = p1.Write(wbuf)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("It will block here because pipe has no buffer!")

	wbuf = []byte("Hello again again!")
	_, err = p2.Write(wbuf)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = p1.Read(rbuf)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(rbuf))

}
