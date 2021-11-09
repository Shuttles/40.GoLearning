package main

import "fmt"

func main()  {
    /*
    append(),向切片的末尾追加数据
        slice = append(slice, elem1, elem2)
        slice = append(slice, anotherSlice...)

     */
    s1 :=[] int{1,2,3,4}
    s2 :=[] int{5,6,7,8}
    //方式一：
    //for i:=0;i<len(s2);i++{
    //    s1 = append(s1,s2[i])
    //}
    // 方法二：
    s1 = append(s1,s2...)
    fmt.Println(s1)

    fmt.Println("---------------")
    //删除切片中的内容
    s3 := []int{1,2,3,4,5,6,7,8,9}
    del:=3//要删除的元素的下标
    //fmt.Println(s3[:del])
    //fmt.Println(s3[del+1:])
    s3 = append(s3[:del],s3[del+1:]...)
    fmt.Println(s3)

}
