package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting the server")
	listener,err:=net.Listen("tcp","127.0.0.1:50000")
	if err!=nil {
		panic(err)
	}
	for{
		conn,err:=listener.Accept()
		if err!=nil {
			fmt.Println("error acceptin",err.Error())
		}
		go doServerConnect(conn)
	}
}

func doServerConnect(conn net.Conn) {
	for  {
		buf:=make([]byte,512)
		len,err:=conn.Read(buf)
		if err!=nil {
			fmt.Println("err reading ",err.Error())
			return
		}
		fmt.Printf("received data: %v",string(buf[:len]))
	}
}
