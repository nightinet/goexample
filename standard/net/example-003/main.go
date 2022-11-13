package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	network := "tcp"
	service := "https"

	if len(os.Args) > 1 {
		network = os.Args[1]

		if len(os.Args) > 2 {
			service = os.Args[2]
		}
	}

	port, err := net.LookupPort(network, service)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Port of [%s:%s] is %d\n", network, service, port)
}
