package main

import (
	"flag"
	"fmt"
)

func main(){
	var param int64
	flag.Int64Var(&param,"uid",1,"uid");
	flag.Parse()
	fmt.Println("param:",param)
	result:=param%16
	fmt.Println("result:",result)
}
