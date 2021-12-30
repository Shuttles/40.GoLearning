package main

import "fmt"

func main() {
	fun1() //调用函数
	fun1()
	fun2 := fun1
	fun2() //调用函数

	//匿名函数：
	func() {
		fmt.Println("我是一个匿名函数。。")
	}()

	fun3 := func() {
		fmt.Println("我也是一个匿名函数。。")
	}

	fun3()
	fun3()

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//定义带返回值的函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 20) //匿名函数调用了，将调用执行的结果给res1
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	} //将匿名函数的值，赋值给res2
	fmt.Println(res2)
}

func fun1() {
	fmt.Println("我是fun1函数。。")
}
