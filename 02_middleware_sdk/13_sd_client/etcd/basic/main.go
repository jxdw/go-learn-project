package main

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"log"
)
func main(){
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.128.128:2379", "192.168.128.128:2479", "192.168.128.128:2579"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	log.Print("etcd connect success")
	defer cli.Close()
	ctx, cancel:=context.WithTimeout(context.Background(),time.Second*2)
	_, err = cli.Put(ctx, "/logagent2/conf7", "sample_value7")
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	resp, err := cli.Get(ctx, "/logagent2/conf7")
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
