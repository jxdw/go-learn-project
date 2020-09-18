package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

func main() {
	nacosTest2()
}

// 我通过example的源码 创建一个真正的注册中心
func nacosTest2()  {
	client, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": []constant.ServerConfig{
			{
				IpAddr: "192.168.172.9",
				Port:   8848,
			},
		},
		"clientConfig": constant.ClientConfig{
			TimeoutMs:           5000,
			NotLoadCacheAtStart: true,
			LogDir:              "data\\server\\nacos\\log",
			NamespaceId: "0cf7435e-262c-4827-86d2-474408a0dced",
			//Username:			 "nacos",
			//Password:			 "nacos",
		},
	})

	if err != nil {
		panic(err)
	}
	param:=vo.RegisterInstanceParam{
		Ip:          "192.168.172.9",
		Port:        8848,
		ServiceName: "demo2.go",
		Weight:      10,
		ClusterName: "zwt",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	}
	success,err:=client.RegisterInstance(param)
	if !success {
		fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, err)
		return
	}
	service,_:=client.GetService(vo.GetServiceParam{
		Clusters: []string{
			"zwt",
		},
		ServiceName: "demo.go",
	})
	fmt.Println("service is ",service)
	time.Sleep(time.Second*10000)
	return
}

