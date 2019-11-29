package main

import (
	"fmt"
	"go-framework-01/golang02_yaml_basic/basic/conf"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println(os.Args[1])
	conf:=new(conf.Yaml)
	yamlFile,err:=ioutil.ReadFile(os.Args[1]+"/test.yaml")
	if err!=nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}
	err=yaml.Unmarshal(yamlFile,conf)
	if err!=nil {
		log.Printf("yaml Unmarshal err %v ", err)
	}
	log.Println("conf",*conf)

}
