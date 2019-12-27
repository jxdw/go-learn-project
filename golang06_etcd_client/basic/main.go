package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	client, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.1.101:2379", "192.168.1.102:2379", "192.168.1.103:2379"},
		DialTimeout: 10 * time.Second,
	})
	defer client.Close()
	kv:=clientv3.NewKV(client);
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	resp, err := kv.Get(context.TODO(), "/prometheus/job",clientv3.WithPrefix())
	///cancel()
	if err != nil {
		println(err)
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
