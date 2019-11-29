package main

import (
	"fmt"
	"github.com/robfig/cron"
	"os"
	"os/signal"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("Every hour on the half hour") })
	go c.Start()
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
