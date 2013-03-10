package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: [discard|echo]", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	if name == "discard" {
		tcpAddr, err := net.ResolveTCPAddr("ip4", ":9000")
		checkError(err)

		listener, err := net.ListenTCP("tcp", tcpAddr)
		checkError(err)

		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go handleDiscard(conn)
		}
	}

	os.Exit(0)
}

func handleDiscard(conn net.Conn) {
	defer conn.Close()
	var b [512]byte
	for {
		conn.Read(b[0:])
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
