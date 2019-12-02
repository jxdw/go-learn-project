package main
import (
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"strconv"
	"time"
)
var host=[]string{
	"http://192.168.172.5:9200/",
	"http://192.168.172.6:9200/",
	"http://192.168.172.7:9200/",
}
var client *elastic.Client
func init(){
	var err  error
	client, err = elastic.NewClient(elastic.SetURL(host...))
	if err != nil {
		fmt.Printf("create client failed, err: %v", err)

	}
}
func PingNode() {
	start := time.Now()

	info, code, err := client.Ping(host[0]).Do(context.Background())
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
	for x<10{
		indexName:="logstash-bbs-2019.11.0"+strconv.Itoa(x);
		fmt.Println(DelIndex(indexName))
		x=x+1
	}
}