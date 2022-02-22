package main

import (
	"fmt"
	"net"
	"os/exec"
)

func main() {
	domain := "dc.services.visualstudio.com"

	cmd := exec.Command("dig", domain)
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("dig ERROR: %s\n", err.Error())
	} else {
		fmt.Printf("dig out: %s\n", string(out))
	}

	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	for _, ip := range ips {
		fmt.Printf("%s. IN A %s\n", domain, ip.String())
	}
}
