package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Printf("usage: %s <host>:<port>\n", args[0])

		return
	}

	connStr := args[1]
	addr, err := net.ResolveUDPAddr("udp4", connStr)
	if err != nil {
		fmt.Printf("error when creating addr: %v\n", err)

		return
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		fmt.Printf("error when connecting: %v\n", err)

		return
	}

	fmt.Printf("Will send data to %s\n", conn.RemoteAddr().String())
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error when reading from stdio: %v\n", err)
		}

		data := []byte(text + "\n")
		if _, err = conn.Write(data); err != nil {
			fmt.Printf("error when writing to conn: %v\n", err)
		}

		if strings.TrimSpace(string(text)) == "QUIT" {
			fmt.Println("UDP client exiting...")

			return
		}

		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Printf("error when reading from conn: %v\n", err)
		}

		fmt.Printf("-> %s", string(buf[0:n]))
	}
}
