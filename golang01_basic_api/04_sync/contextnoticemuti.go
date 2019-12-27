package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	cancelContext,cancel:=context.WithCancel(context.Background())
	go watch(cancelContext,"监控1")
	go watch(cancelContext,"监控2")
	go watch(cancelContext,"监控3")
	time.Sleep(10*time.Second)
	fmt.Println("可以了，通知监控退出")
	cancel()
	time.Sleep(5*time.Second)
}
func watch(ctx context.Context,name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了")
			return
		default:
			fmt.Println(name, "goroutine监控中")
			time.Sleep(1 * time.Second)
		}
	}
}