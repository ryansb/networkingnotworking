package main

import (
    "net"
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Fprintf(os.Stderr, "Usage: [discard|echo]", os.Args[0])
        os.Exit(1)
    }
    name := os.Args[1]

    if name == "discard" {
        tcpAddr, err := net.ResolveTCPAddr("ip4", ":9")
        checkError(err)

        listener, err := net.ListenTCP("tcp", tcpAddr)
        checkError(err)

        var b [512]byte
        for {
            conn, err := listener.Accept()
            checkError(err)
            conn.Read(b[0:])
        }
    }

    os.Exit(0)
}


func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}
