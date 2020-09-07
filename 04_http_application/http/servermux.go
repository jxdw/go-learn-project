package main

import "net/http"

func main(){
	mux:=http.NewServeMux();
	mux.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("enjoy:grpc sever"))
	})
	http.ListenAndServe(":12345",mux)
}
