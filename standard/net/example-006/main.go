package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatalln(err)
	}

	for _, addr := range addrs {
		fmt.Println(addr.Network(), addr.String())
	}
}
