package main

import (
	"fmt"
	dis "go-learn-project/02_middleware_sdk/06_etcd_client/servicediscover/discovery"
	"log"
	"time"
)

func main() {

	m, err := dis.NewMaster([]string{
		"http://127.0.0.1:2379",
	}, "/services/")

	if err != nil {
		log.Fatal(err)
	}

	for {
		for k, v := range  m.Nodes {
			fmt.Printf("node:%s, ip=%s\n", k, v.Info.IpAddress)
		}
		fmt.Printf("nodes num = %d\n",len(m.Nodes))
		time.Sleep(time.Second * 5)
	}
}
