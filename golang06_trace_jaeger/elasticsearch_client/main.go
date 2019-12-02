package main

import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)
type Conf struct {
	IndexPrefix string `yaml:"indexprefix"`
	HostAddress []string `yaml:"hostAddress"`
	MaxInt  int `yaml:"maxInt"`
}
var conf Conf
var client *elastic.Client

func init(){
	conf=Conf{}
	yamlFile,err:=ioutil.ReadFile(os.Args[1]+"/conf.yaml")
	if err!=nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err=yaml.Unmarshal(yamlFile,&conf)
	if err!=nil {
		log.Printf("yaml Unmarshal err %v ", err)
	}

	client, err = elastic.NewClient(elastic.SetURL(conf.HostAddress...))
	if err != nil {
		fmt.Printf("create client failed, err: %v", err)

	}
}
func PingNode() {
	start := time.Now()

	info, code, err := client.Ping(conf.HostAddress[0]).Do(context.Background())
	if err != nil {
		fmt.Printf("ping es failed, err: %v", err)
	}

	duration := time.Since(start)
	fmt.Printf("cost time: %v\n", duration)
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

//删除 index
func DelIndex(index... string) bool {
	response, err := client.DeleteIndex(index...).Do(context.Background())
	if err != nil {
		fmt.Printf("delete index failed, err: %v\n", err)
	}
	return response.Acknowledged
}
func main() {
	PingNode()
	//fmt.Println(time.Now().Format("2018.01.01"))
	var x int=1
	for x < conf.MaxInt {
		var indexName string
		if x<10 {
			indexName=conf.IndexPrefix+"0"+strconv.Itoa(x);
		}else {
			indexName=conf.IndexPrefix+strconv.Itoa(x);
		}
		fmt.Println(DelIndex(indexName))
		x=x+1
	}
}