package main

import "fmt"

func main() {
	/*

	   上节课回顾：
	       A：回调函数
	           一个函数fun1,可以接受另一个函数fun2作为参数。
	               高阶函数：fun1
	               回调函数(callback)：fun2

	*/

	res1 := increment()      //res1 = fun
	fmt.Printf("%T\n", res1) //func() int
	fmt.Println(res1)
	v1 := res1()
	fmt.Println(v1) //1
	v2 := res1()
	fmt.Println(v2)     //2
	fmt.Println(res1()) //3
	fmt.Println(res1()) //4
	fmt.Println(res1()) //5
	fmt.Println(res1()) //6

	res2 := increment()
	v3 := res2() //1
	fmt.Println(v3)
	v4 := res2() //2
	fmt.Println(v4)

	v5 := res1()
	fmt.Println(v5) //7

}

//可以将一个函数作为返回值。
func increment() func() int { //外层函数
	//1.定义了一个局部变量
	i := 0

	//2.定义了一个匿名函数，给变量自赠了并返回
	fun := func() int { //内层函数
		i++
		return i
	}
	//3.返回该匿名函数
	return fun
}
