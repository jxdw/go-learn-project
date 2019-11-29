package main

import ("fmt"
	"github.com/robfig/cron"
	"os"
	"os/signal"
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
	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("Every hour on the half hour") })
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
