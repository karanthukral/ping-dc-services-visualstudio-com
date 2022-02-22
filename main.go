package main

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"time"
)

func main() {
	domain := "dc.services.visualstudio.com"

	cmd := exec.Command("dig", domain, "@1.1.1.1")
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("dig ERROR: %s\n", err.Error())
	} else {
		fmt.Printf("dig out: %s\n", string(out))
	}

	resolver := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return d.DialContext(ctx, "udp", "1.1.1.1:53")
		},
	}

	addrs, err := resolver.LookupHost(context.Background(), domain)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	for _, ip := range addrs {
		fmt.Printf("%s. IN %s\n", domain, ip)
	}
}
