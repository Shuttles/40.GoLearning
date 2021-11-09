package main

import "fmt"

type Dog struct {
    age int
}

type Cat struct{
    weigh float64
}

type Animal1 interface {

}

//使用空接口，接收任意类型作为参数
func info(v interface{})  {
    fmt.Println(v)
}


func main()  {
    /*
    使用空接口，可以实现各种类型的对象存储。
     */
    d1:= Dog{1}
    d2 := Dog{2}
    c1 :=Cat{3.2}
    c2:=Cat{3.5}

    animals:=[4] Animal1{d1,d2,c1,c2}
    fmt.Println(animals)


    info(d1)
    info(c1)
    info("aaa")
    info(100)
}
