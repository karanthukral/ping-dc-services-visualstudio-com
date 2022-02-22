package main

import (
	"fmt"
	"net"
)

func main() {
	domain := "dc.services.visualstudio.com"
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	for _, ip := range ips {
		fmt.Printf("%s. IN A %s\n", domain, ip.String())
	}
}
