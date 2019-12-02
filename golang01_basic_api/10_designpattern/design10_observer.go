package main

import "fmt"

type Observer interface {
	Update(subject *Subject)
}

type Subject struct {
	observers []Observer
	context string
}
func (s *Subject) notify() {
	for _, o := range s.observers {
		o.Update(s)
	}
}
func (s *Subject)Attach(o Observer){
	s.observers = append(s.observers, o)
}
func (s *Subject) UpdateContext(context string) {
	s.context = context
	s.notify()
}
func NewSubject() *Subject {
	return &Subject{observers:make([]Observer,0)}
}
type  Reader struct {
	name  string
}
func (r *Reader) Update(s *Subject) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
func NewReader(name string) *Reader{
	return &Reader{name:name}
}
func main() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)
	subject.UpdateContext("observer modee")
}
