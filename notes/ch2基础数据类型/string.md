# 1.string介绍

1. 一个字符串是一个==不可改变==的**字节序列**。字符串可以包含**任意的数据**，包括byte值0，但是通常是用来包含人类可读的文本。文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。

   Go中的字符串是Unicode兼容的，并且是UTF-8编码的。

   示例代码：

   ```go
   func main() {  
       name := "Hello World"
       fmt.Println(name)
   }
   ```

2. string是只读的！！！不可以修改



# 2.操作

## 2.1len函数

1. 内置的len函数可以返回一个字符串中的==字节数目==（不是rune字符数目），索引操作s[i]返回第i个字节的字节值，i必须满足0 ≤ i< len(s)条件约束。

   ```Go
   s := "hello, world"
   fmt.Println(len(s))     // "12"
   fmt.Println(s[0], s[7]) // "104 119" ('h' and 'w')
   ```

   如果试图访问超出字符串索引范围的字节将会导致**panic异常**：

   ```Go
   c := s[len(s)] // panic: index out of range
   ```

   ==第i个字节并不一定是字符串的第i个字符，因为对于非ASCII字符的UTF8编码会要两个或多个字节。==



## 2.2新串

1. 子字符串操作`s[i:j]`基于原始的s字符串的第i个字节开始到第j个字节（**并不包含j本身**）生成一个新字符串。生成的新字符串将包含**j-i个字节**。

   ```Go
   fmt.Println(s[0:5]) // "hello"
   ```

   同样，如果索引超出字符串范围或者j小于i的话将导致**panic异常**。

   不管i还是j都可能被忽略，当它们被忽略时将采用0作为开始位置，采用len(s)作为结束的位置。

   ```Go
   fmt.Println(s[:5]) // "hello"
   fmt.Println(s[7:]) // "world"
   fmt.Println(s[:])  // "hello, world"
   ```

## 2.3不变性

1. 因为字符串是不可修改的，因此尝试修改字符串内部数据的操作也是被禁止的：

   ```Go
   s[0] = 'L' // compile error: cannot assign to s[0]
   ```

   不变性意味着如果两个字符串**共享相同的底层数据**的话也是安全的，这使得**复制**任何长度的字符串代价是**低廉**的。

   同样，一个字符串s和对应的子字符串切片（==这里的切片并不是切片类型==）s[7:]的操作也可以**安全地共享相同的内存**，因此字符串切片操作代价也是低廉的。在这两种情况下都==没有必要分配新的内存==。 图3.4演示了一个字符串和两个子串共享相同的底层数据。

   ![](https://books.studygolang.com/gopl-zh/images/ch3-04.png)





# string和数字转换

https://books.studygolang.com/gopl-zh/ch3/ch3-05.html



# strings包中的常用方法

https://www.chaindesk.cn/witbook/13/182