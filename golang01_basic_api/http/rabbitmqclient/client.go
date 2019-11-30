package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)
/**
1.代码作用: 访问rabbitmq rest api，得到一些数据。后续可以根据这些数据作监控告警功能。
2.最新的rabbitmq api文档地址: https://rawcdn.githack.com/rabbitmq/rabbitmq-management/v3.8.1/priv/www/api/index.html
3.参考资料: https://supereagle.github.io/2017/11/22/request-with-auth/
 */
func main() {
	result,err:=do("GET","http://192.168.172.3:15672/api/connections",nil);
	if err!=nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	fmt.Println(buf.String())
}
func do(method string,url string,playload io.Reader) (*http.Response,error)  {
	req,err:=http.NewRequest(method,url,playload)
	if err!=nil {
		return nil,err
	}
	req.SetBasicAuth("admin","admin")
	return  http.DefaultClient.Do(req)
}
