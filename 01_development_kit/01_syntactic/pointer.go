package main

import (
	"fmt"

)

func chavalue(a int)  int{
	fmt.Println("chavalue a address is ",&a)
	a=a+1
	return a
}
func changePointer(a *int) {
	*a=*a+2
	return
}
func main() {
	a:=10;
	fmt.Println("a address is ",&a)
	b:=chavalue(a)

	fmt.Println(a)
	fmt.Println(b)
	changePointer(&a)
	fmt.Println(a)
}
