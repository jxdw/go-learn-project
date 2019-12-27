package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
	//context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点
	//使用context.WithCancel(parent)函数，创建一个可取消的子Context,然后当作参数传给goroutine使用
	//使用这个子Context跟踪这个goroutine
	cancelContext,cancel:=context.WithCancel(context.Background());
	go func(ctx context.Context) {
		for  {
			select{
			case <-ctx.Done():
				fmt.Println("监控退出，停止了")
				return
			default:
				fmt.Println("goroutine监控中")
				time.Sleep(1*time.Second)
			}
		}
	}(cancelContext)
	time.Sleep(10*time.Second)
	fmt.Println("通知监控退出")
	cancel()
	time.Sleep(5*time.Second)
}
