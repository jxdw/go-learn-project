package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)
//实现http.HandlerFunc接口，
//函数签名符合http.HandlerFunc
func hello(w http.ResponseWriter,request *http.Request){
	s,_:=ioutil.ReadAll(request.Body)
	body:=string(s)
	fmt.Println("second")

	fmt.Println(body)
	w.Write([]byte("hello world"))
}

func log(hf http.HandlerFunc) http.HandlerFunc{
	count:=0;
	return func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("first")
		s,_:=ioutil.ReadAll(request.Body)
		body:=string(s)
		fmt.Println(body)
		request.Body = ioutil.NopCloser(bytes.NewBuffer(s))
		count++;
		invokeinfo:="handler function  call "+strconv.Itoa(count)+" times"
		fmt.Println(invokeinfo)
		hf(writer,request)
	}
}

func main(){
	server:=http.Server{Addr:"127.0.0.1:9999"}
	http.HandleFunc("/hello",log(hello))
	server.ListenAndServe()
}
