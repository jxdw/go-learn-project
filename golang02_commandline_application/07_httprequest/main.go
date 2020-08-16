package main

import (
"fmt"
"os"
"strings"
"net/http"
"io/ioutil"
)

func main() {
	url := os.Args[1]
	method := "POST"

	payload ,err:= ioutil.ReadFile("post.json")

	client := &http.Client {}
	reqdata:=string(payload)
	fmt.Println("request data is",reqdata)
	req, err := http.NewRequest(method, url, strings.NewReader(reqdata))

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
