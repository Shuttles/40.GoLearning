package main

import "fmt"

type sharp interface {
	area() float64
}

type sqrt struct {
	l float64
}

func (s sqrt) area() float64 {
	return s.l * s.l
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return c.r * 3.14 * c.r
}

// 接口类型作为函数参数
func getArea(s sharp) {
	fmt.Println(s.area())
}

func main() {
	s1 := sqrt{2}
	c1 := circle{1}

	getArea(s1)
	getArea(c1)
}
