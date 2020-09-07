package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	buf:=make([]byte,512);
	var err error
	buf,err=ioutil.ReadFile(os.Args[1]+"/test.txt")
	if err!=nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
