package main

import ("fmt"

)



func chavalue(a int)  int{
	a=a+1
	return a
}
func changePointer(a *int) {
	*a=*a+2
	return
}
func main() {
	a:=10;
	chavalue(a)
	fmt.Println(a)

	changePointer(&a)
	fmt.Println(a)

}
