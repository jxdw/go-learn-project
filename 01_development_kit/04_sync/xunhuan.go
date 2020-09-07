package main

import "fmt"

var messages chan string = make(chan string)

func main() {
	go setMessage()
	for  {
		select {
			default:
				fmt.Println(<-messages)
		} // 取消息
	}
}
func setMessage() {
	for  {
		messages <- "ping" // 存消息
	}
}