package main

import (
	"fmt"
)
func a() int64 {
	fmt.Println("calling a()")
	return 2
}
var T int64 = a()

func init() {
	fmt.Println("init in main.go ")
}


func main() {
	fmt.Println("calling main")
}
