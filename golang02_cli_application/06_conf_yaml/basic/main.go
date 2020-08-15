package main

import (
	"fmt"
	"go-learn-project/golang02_cli_application/06_conf_yaml/basic/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	conf:=new(conf.Config)
	yamlFile,err:=ioutil.ReadFile(os.Args[1]+"/config.yaml")
	if err!=nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err=yaml.Unmarshal(yamlFile,conf)
	if err!=nil {
		log.Printf("yaml Unmarshal err %v ", err)
	}
	log.Println("conf",conf.Mysql)

}
