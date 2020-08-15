package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go func() {
		time.Sleep(time.Microsecond*2)
		fmt.Println("1号工作完成")
		waitGroup.Done()
	}()
	go func() {
		time.Sleep(time.Microsecond*2)
		fmt.Println("2号工作完成")
		waitGroup.Done()
	}()
	waitGroup.Wait()
	fmt.Println("好了，大家活都干完了，收工")
}
