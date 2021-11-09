package main

import "fmt"

func main()  {
    /*
    copy(),
        切片是引用类型，传递的是地址

    深拷贝和浅拷贝
        深拷贝：拷贝的是数据
            值类型都是深拷贝，基本类型,数组

        浅拷贝：拷贝的是地址
            引用类型默认都是浅拷贝，切片，map
     */
     s1 := []int{1,2,3,4,5}
     s2 := s1
     fmt.Println(s1,s2)
     s2[0] =100
     fmt.Println(s1,s2)

     a := 100
     b := a
     fmt.Println(a,b)
     b = 200
     fmt.Println(a,b)
     fmt.Println("-------------")

     m:=[]int{1,2,3,4,5}
     n:=[]int{7,8,9}
     fmt.Println(m)
     fmt.Println(n)

     //copy(m,n) //将n中的数据，拷贝到m里
     //copy(n,m)//将m中的数据，拷贝到n里
     //copy(n,m[1:4])//将m中的下标1到3的数据，拷贝到n里
     copy(n[1:],m[3:]) //将m中的下标3到最后的数据，拷贝到n的下标1之后
     fmt.Println(m)
     fmt.Println(n)


     s3 :=[]int{1,2,3,4}
     s4 := s3 //浅拷贝

     s5:=make([]int,4,4)
     copy(s5,s3)//深拷贝
     fmt.Println(s4)
     fmt.Println(s5)
     s3[0] = 100
     fmt.Println(s4)
     fmt.Println(s5)

}
