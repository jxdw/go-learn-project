package main

import "net/http"

func main()  {
	http.HandleFunc("/api/{serviceName}",func( response http.ResponseWriter,request *http.Request) {

	})
	http.ListenAndServe(":8080",nil)
}
