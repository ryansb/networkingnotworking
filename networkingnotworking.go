package main

import (
	"fmt"
	"net"
	"os"
)
var asciichars = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"


func main() {
	go echoServer(":8007")
	go discardServer(":8009")
	chargenServer(":8019")
}

func echoServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("ip4", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleEcho(conn)
	}
}

func discardServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("ip4", port)
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

func chargenServer(port string) {
	tcpAddr, err := net.ResolveTCPAddr("ip4", port)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleChargen(conn)
	}
}

func handleDiscard(conn net.Conn) {
	defer conn.Close()
	var b [512]byte
	for {
		conn.Read(b[0:])
	}
}

func handleChargen(conn net.Conn) {
	var b [80]byte
	copy(b[:], asciichars)
	defer conn.Close()
	conn.Write(b[0:72])

	var n [1]byte
	copy(n[:], "\n")
	conn.Write(n[:])
}

func handleEcho(conn net.Conn) {
	defer conn.Close()
	var b [512]byte

	n, err := conn.Read(b[0:])

	if err != nil {
		return
	}

	_, err2 := conn.Write(b[0:n])
	if err2 != nil {
		return
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
