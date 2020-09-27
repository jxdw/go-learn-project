package main

import (
	"fmt"
	"time"
)

type Task struct {
	fc func() error
}

func (t *Task) Execute()  {
	t.fc()
}

func NewTask(f func() error) *Task {
	t:=Task{
		fc:f,
	}
	return &t
}

type Pool struct {
	EntryChannel chan *Task
	work_num int
	JobsChannel chan *Task
}
func (p *Pool) work(i int) {
	for task:=range p.JobsChannel {
		task.Execute()
		fmt.Println("work ID",i," work done")
	}

}

func (p *Pool) Run(){
 	for i:=0;i<p.work_num;i++{
 		go p.work(i)
	}
	for task:= range p.EntryChannel {
		p.JobsChannel<-task
	}
	close(p.JobsChannel)
	close(p.EntryChannel)
 }

func NewPool(cap int) *Pool{
	p:=Pool{
		EntryChannel: make(chan *Task),
		work_num: cap,
		JobsChannel: make(chan *Task),
	}
	return &p
}

func main()  {
	t:=NewTask(func() error{
		flowOut:="2006-01-02 15:04:05.000 "
		fmt.Print(time.Now().Format(flowOut))
		return nil
	})
	p:=NewPool(3)
	go func() {
		for {
			p.EntryChannel<-t
		}
	}()
	p.Run()
}