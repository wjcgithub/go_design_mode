package main

import "fmt"

type Observer interface {
	Update()
}

type Subject interface {
	Attach(o Observer)
	Detach(o Observer)
	Notify()
}

type Oberver1 struct {
}

func (o *Oberver1) Update() {
	fmt.Println("Oberver1 update")
}

type Oberver2 struct {
}

func (o Oberver2) Update() {
	fmt.Println("Oberver2 update")
}

type SubjectImpl struct {
	observers []Observer
}

func (s *SubjectImpl) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *SubjectImpl) Detach(o Observer) {
	for i, v := range s.observers {
		if v == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

func (s *SubjectImpl) Notify() {
	for _, observer := range s.observers {
		observer.Update()
	}
}

func main() {
	subject := &SubjectImpl{}
	observer1 := &Oberver1{}
	observer2 := Oberver2{}
	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Notify()
}
