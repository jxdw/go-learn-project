package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

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
