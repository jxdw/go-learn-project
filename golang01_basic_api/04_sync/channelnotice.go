package main

import (
	"fmt"
	"time"
)

func main(){
	stopchan :=make(chan bool)
	go func() {
		for  {
			select {
				case <-stopchan:
					fmt.Println("监控退出了，停止了")
					return
				default:
					fmt.Println("goroutine监控中...")
					time.Sleep(time.Second*1)
			}
		}
	}()
	time.Sleep(10*time.Second);
	fmt.Println("可以了通知监控退出")
	stopchan <- true
	time.Sleep(5*time.Second)
}
