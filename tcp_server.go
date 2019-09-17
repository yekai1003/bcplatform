package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func handle_conn(conn net.Conn) {
	for {
		n, err := io.Copy(os.Stdout, conn)
		if err != nil {
			fmt.Println("Failed to Copy ", err)
			break
		}
		if n <= 0 {
			fmt.Println("Client closed ")
			break
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to Listen ", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept ", err)
			continue
		}
		go handle_conn(conn)
	}
}
