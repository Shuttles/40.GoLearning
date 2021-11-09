package main

import "fmt"

func main() {
    s := "Hello world"
    fmt.Println(s)

    s1 := s[:5]
    fmt.Println(s1)

    s1[0] = 'h' // error : 不可以修改字符串
    fmt.Println(s1)
}

