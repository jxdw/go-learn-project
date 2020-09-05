package main

import (
	"fmt"
	"time"
)

func main() {
	timer:=time.NewTimer(time.Duration(time.Second*2))
	defer timer.Stop()
	for  {
		<-timer.C
		fmt.Println("timeout...")
		timer.Reset(time.Second*2)

	}
}
