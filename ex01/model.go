package ex01

import "fmt"

type Human struct {
	age  uint8
	name string
}

func NewHuman(age uint8, name string) *Human {
	return &Human{age: age, name: name}
}

func (h *Human) SayHello() {
	fmt.Printf("Hello, I'm %s, I'm %d\n", h.name, h.age)
}

//Встраивание (Композиция) Human в Action
type Action struct {
	Human
	Idea string
}
