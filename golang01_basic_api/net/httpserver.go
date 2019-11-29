package main

import "net/http"

func main()  {
	http.HandleFunc("/api/serviceName",func( response http.ResponseWriter,request *http.Request) {
       response.Write([]byte("hello word"))
	})
	http.ListenAndServe(":8080",nil)
}
