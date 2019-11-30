package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn,err:=net.Dial("tcp","127.0.0.1:50000")
	if err!=nil {

	}
	inputReader:=bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name?")
	clientName,_:=inputReader.ReadString('\n')
	trimmedClient:=strings.Trim(clientName,"\r\n")
	for{
		fmt.Println("what to send  to the server,type q to quit")
		input,_:=inputReader.ReadString('\n')
		trimmedInput:=strings.Trim(input,"\r\n")
		if trimmedInput=="Q" {
			return
		}
		_,err=conn.Write([]byte(trimmedClient +" says: "+ trimmedInput))
	}
}
