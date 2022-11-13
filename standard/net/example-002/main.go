package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	host := "www.google.com"

	if len(os.Args) > 1 {
		host = os.Args[1]
	}

	fmt.Printf("IP of [%s] are:\n", host)
	addrs, err := net.LookupHost(host)
	if err != nil {
		log.Fatalln(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr)
	}
	fmt.Println()

	for {
		cname, err := net.LookupCNAME(host)
		if err != nil {
			log.Fatalln(err)
		} else if cname != host && cname != host+"." {
			fmt.Printf("CNAME of [%s] is: [%s]\n", host, cname)
			host = cname
		} else {
			break
		}
	}

	for _, addr := range addrs {
		fmt.Println()

		hosts, err := net.LookupAddr(addr)
		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Printf("Name of [%s] are:\n", addr)
		for _, host := range hosts {
			fmt.Println(host)
		}
	}
}
