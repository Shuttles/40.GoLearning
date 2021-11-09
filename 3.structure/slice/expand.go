package main

import (
    "fmt"
)

func main()  {
/*
    切片叫做变长数组，长度不固定
        len（），cap()

    当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容(成倍增长)
    3-->6-->12-->24-->48
    4-->8-->16-->32
     */
    arr := [10]int{1,2,3,4,5,6,7,8,9,10}
    fmt.Println(arr)
    fmt.Printf("%p\n",&arr) //打印输出数组自己的地址

    s1 := arr[:5]
    fmt.Println(s1)
    fmt.Printf("%p,长度：%d，容量：%d\n",s1,len(s1),cap(s1))
    s2 := arr[2:7]
    fmt.Println(s2)
    fmt.Printf("%p,长度：%d，容量：%d\n",s2,len(s2),cap(s2))

    //修改切片的数据,其实是修改该切片所指向的底层数组的数据。
    //导致所有指向该数组的切片的数据，都会改变。
    s1[3] = 100
    fmt.Println(arr)
    fmt.Println(s1)
    fmt.Println(s2)

    //append(),追加数据
    s1 = append(s1,1,1,1,1) //s1的长度：5，容量：10
    fmt.Println(arr) //[1 2 3 100 5 1 1 1 1 10]
    fmt.Println(s1) //[1 2 3 100 5 1 1 1 1]
    fmt.Println(s2) //[3 100 5 1 1]

    //追加的时候，涉及到了扩容,会更改切片指向的底层数组
    //s2:len5,cap8
    s2= append(s2, 2,2,2,2,2)
    //更改s2指向的底层数组
    fmt.Println(arr)//[1 2 3 100 5 1 1 1 1 10]
    fmt.Println(s1) //[1 2 3 100 5 1 1 1 1]
    fmt.Println(s2) //[3 100 5 1 1 2 2 2 2 2]
    fmt.Printf("%p,长度：%d，容量：%d\n",s2,len(s2),cap(s2))

}
