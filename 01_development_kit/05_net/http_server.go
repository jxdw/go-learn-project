package main

import (
	"fmt"
	"net/http"
)

type Helloword struct {

}
func (h *Helloword) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c,err:=r.Cookie("youxia");
	if err!=nil {
		fmt.Println(err)
	}else {
		fmt.Printf("%s=%s",c.Name,c.Value)
	}
	w.Write([]byte("Hello, world!"))
}
func main()  {
	http.Handle("/", &Helloword{})
	http.ListenAndServe("10.3.20.215:8888", nil)
}
