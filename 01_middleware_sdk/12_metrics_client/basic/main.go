package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-learn-project/01_middleware_sdk/12_metrics_client/basic/config"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	conf:=new(config.AppConf)
	yamlFile,err:=ioutil.ReadFile(os.Args[1]+"/config.yaml")
	if err!=nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err=yaml.Unmarshal(yamlFile,conf)
	http.Handle("/metrics",promhttp.Handler())
	http.ListenAndServe(conf.Httpaddress, nil)
}
