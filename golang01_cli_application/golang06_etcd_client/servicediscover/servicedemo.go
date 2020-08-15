package main

import (
	"fmt"
	dis "go-learn-project/golang01_cli_application/golang06_etcd_client/servicediscover/discovery"
	"log"
	"time"
)

func main() {

	serviceName := "s-test"
	serviceInfo := dis.ServiceInfo{IpAddress: "192.168.1.105"}

	s, err := dis.NewService(serviceName, serviceInfo,[]string {
		"http://127.0.0.1:2379",
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name:%s, ip:%s\n", s.Name, s.Info.IpAddress)


	go func() {
		time.Sleep(time.Second*600)
		s.Stop()
	}()

	s.Start()
}
