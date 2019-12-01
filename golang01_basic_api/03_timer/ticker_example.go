package main

import (
	"fmt"
	"time"
)

func main() {
	t:=time.NewTicker(time.Duration(time.Second*2))
	defer t.Stop()

	for  {
		<-t.C
		fmt.Println("timeout...")
	}
}
