package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	peri() float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	p := t.peri() / 2
	area := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return area
}

type Circle struct {
	r float64
}

func (c Circle) peri() float64 {
	return 2 * math.Pi * c.r
}
func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// test
func testArea(s Shape) {
	fmt.Println("面积： ", s.area())
}

func testPeri(s Shape) {
	fmt.Printf("周长：%.2f\n", s.peri())
}

func getType(s Shape) {
	if inst, ok := s.(Triangle); ok {
		fmt.Println("是Triangle类型。。三边是：", inst.a, inst.b, inst.c)
	} else if inst, ok := s.(Circle); ok {
		fmt.Println("是Circle类型，半径是：", inst.r)
	} else {
		fmt.Println("以上都不对。。")
	}
}

func getType2(s Shape) {
	switch inst := s.(type) {
	case Triangle:
		fmt.Println("三角形啊。。", inst.a, inst.b, inst.c)
	case Circle:
		fmt.Println("圆形啊。。", inst.r)
	}
}

func main() {
	t := Triangle{3, 4, 5}
	testArea(t)

	c := Circle{1}
	testPeri(c)

	//定义一个接口类型的数组：Shape类型，可以存储该接口的任意实现类的对象作为数据。
	var arr [4]Shape
	arr[0] = t
	arr[1] = c
	arr[2] = Triangle{1, 2, 3}
	arr[3] = Circle{2}

	// 判断类型
	getType(t)
	getType2(c)
}
