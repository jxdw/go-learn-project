package main

import "time"
type T struct{
	H int
	I int
	L int
	J int
	K int
	M int
	N int
}
type value struct {
	A string
	B int
	C time.Time
	D  []byte
	E float32
	F *string
	T T
}

func main()  {
	m:=make(map[int]*value,10000000)
	for i:=0;i<10000000;i++ {
		m[i]=&value{}
	}
	for i:=0;;i++ {
		delete(m,i)
		m[10000000+i]=&value{}
		time.Sleep(time.Second*5)
	}
}