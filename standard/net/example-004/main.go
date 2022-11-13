package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	cidr := "192.168.0.0/24"

	if len(os.Args) > 1 {
		cidr = os.Args[1]
	}

	addr, net, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(addr, net)
}
