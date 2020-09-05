package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	//生成可被取消的上下文
	cancelContext,cancel:=context.WithDeadline(context.Background(),time.Now().Add(8 * time.Second))
	go watch(cancelContext,"监控1")
	go watch(cancelContext,"监控2")
	go watch(cancelContext,"监控3")
	time.Sleep(10*time.Second)
	fmt.Println("可以了，通知监控退出")
	//取消上下文
	cancel()
	time.Sleep(5*time.Second)
}
func watch(ctx context.Context,name string) {
	for {
		select {
		//上下文是否被取消
		case <-ctx.Done():
			fmt.Println(name, "监控退出，停止了")
			return
		default:
			//必须是这个时间点, 据说是go诞生之日
			fmt.Println(time.Now().Format("2006-01-02 15:04:05")," ",name, "goroutine监控中")
			time.Sleep(1 * time.Second)
		}
	}
}