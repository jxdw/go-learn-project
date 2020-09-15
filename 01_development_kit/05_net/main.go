package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	lsr, err := net.Listen("tcp", ":7070")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		return
	}

	for {
		conn , err := lsr.Accept()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			continue
		}

		go connHandler(conn)

	}

	fmt.Println("Done !")
}

func connHandler(conn net.Conn) {
	defer conn.Close()

	var buf[512]byte
	for {
		n , err := conn.Read(buf[0:])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
		_, err = conn.Write(buf[0:n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
			return
		}
	}
}