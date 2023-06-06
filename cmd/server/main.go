package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: %s <port>\n", args[0])

		return
	}

	port := ":" + args[1]
	addrStr, err := net.ResolveUDPAddr("udp4", port)
	if err != nil {
		fmt.Printf("error when creating addr: %v\n", err)

		return
	}
	conn, err := net.ListenUDP("udp4", addrStr)
	if err != nil {
		fmt.Printf("error creating listener: %v\n", err)

		return
	}
	defer conn.Close()

	fmt.Printf("Listening on %s\n", port)

	buf := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("error reading from conn: %v\n", err)

			continue
		}
		if strings.TrimSpace(string(buf[0:n])) == "QUIT" {
			fmt.Println("Exiting UDP server...")

			return
		}

		fmt.Print("-> ", string(buf[0:n-1]))
		t := time.Now()
		tf := t.Format(time.RFC3339) + "\n"

		if _, err = conn.WriteToUDP([]byte(tf), addr); err != nil {
			fmt.Printf("error reading from conn: %v\n", err)
		}
	}
}
