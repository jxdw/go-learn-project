package main

import (
	"fmt"
)

type Api interface {
	Say(name string) string
}

type  hiApi struct {

}
func (*hiApi) Say(name string) string{
	return fmt.Sprintf("hi,%s",name)
}

type HelloApi struct {

}
func (*HelloApi) Say(name string) string{
	return fmt.Sprintf("hello,%s",name)
}

func NewApi(target int) Api{
	if target==1{
		return &hiApi{}
	}
	if target==2 {
		return &HelloApi{}
	}
	return nil
}
func main() {
	api := NewApi(1)
	s := api.Say("Tom")
	fmt.Println(s)
	if s != "Hi, Tom" {
		fmt.Errorf("Type1 test fail")
	}
	api=NewApi(2)
	s=api.Say("Tom2")
	fmt.Println(s)
	if s != "Hi, Tom2" {
		fmt.Errorf("Type2 test fail")
	}
}
