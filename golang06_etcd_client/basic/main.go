package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)
//client的v3在new{clientv3.Config{}}
//client的v3版本提供六类api：
//KV          -clientv3.NewKV(client)、K-V键值库操作
//cluster     -clientv3.NewCluster(client)、向集群里增加删除etcd节点，集群管理(用的比较少)
//auth        -clientv3.AuthEnable(client)、管理etcd的用户和权限，属于管理员操作(用的比较少)
//watch       -clientv3.NewWatch(client)、观察订阅，从而监听最新的数据变化
//lease       -clientv3.NewLease(client)、租约相关操作,租约过期会删除授权的key
//Maintenance -clientv3.NewMaintenance(client)、维护etcd，比如主动迁移etcd的leader节点(用的比较少)
func main() {
	client, _ := clientv3.New(clientv3.Config{
		Endpoints:   []string{"192.168.172.5:2379"},
		DialTimeout: 10 * time.Second,
	})
	defer client.Close()
	cluster:=clientv3.NewCluster(client);
	memberListResponse,err:=cluster.MemberList(context.TODO())
	fmt.Printf("etcd member list:\n")
	for _,member:=range memberListResponse.Members{
		fmt.Printf("%s : %s\n",member.Name,member.PeerURLs)
	}
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
