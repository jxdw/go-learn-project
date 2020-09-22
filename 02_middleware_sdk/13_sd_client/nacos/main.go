package main

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

//创建一个真正的注册中心
func init()  {
	serverconfig:=[]constant.ServerConfig{
		{
			IpAddr: "192.168.172.9",
			Port:   8848,
		},
	}
	clientconfig:=constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "\\data\\server\\nacos\\log",
		NamespaceId: "0cf7435e-262c-4827-86d2-474408a0dced",
		CacheDir: "\\data\\server\\demoservice",
		//Username:			 "nacos",
		//Password:			 "nacos",
	}

	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs":serverconfig ,
		"clientConfig": clientconfig,
	})

	if err != nil {
		panic(err)
	}

	param:=vo.RegisterInstanceParam{
		Ip:          "10.3.20.215",
		Port:        18848,
		ServiceName: "go_demo_service",
		Weight:      10,
		ClusterName: "zwt",
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	}
	success,err:=namingClient.RegisterInstance(param)

	if !success {
		fmt.Printf("RegisterServiceInstance,param:%+v,result:%+v \n\n", param, err)
		return
	}

	service,_:=namingClient.GetService(vo.GetServiceParam{
		Clusters: []string{
			"zwt",
		},
		ServiceName: "go_demo_service",
	})
	fmt.Println("service is ",service)


	configclient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverconfig,
		"clientConfig":  clientconfig,
	})


	configFile,err:=configclient.GetConfig(vo.ConfigParam{
		DataId: "test-data",
		Group:  "test-group",
	})

	fmt.Println("config file  content is :"+configFile)

	for success {
		err = configclient.ListenConfig(vo.ConfigParam{
			DataId: "test-data",
			Group:  "test-group",
			OnChange: func(namespace, group, dataId, data string) {
				fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
			},
		})
		if err!=nil {
			fmt.Println("error is ",err)
		}
	}

}

func main() {

}
