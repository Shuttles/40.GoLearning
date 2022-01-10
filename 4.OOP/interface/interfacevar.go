package main

import "fmt"

type Men interface {
	SayHi()
	Sing(string)
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(song string) {
	fmt.Println("La la la la...", song)
}
