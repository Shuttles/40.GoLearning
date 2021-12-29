# 1.操作

1. **定义和初始化**

   ```go
   type struct_variable_type struct {
      member definition;
      member definition;
      ...
      member definition;
   }
   ```

   一旦定义了结构体类型，它就能用于变量的声明

   ```go
   variable_name := structure_variable_type {value1, value2...valuen}
   ```

   **初始化结构体**

   ```go
   // 1.按照顺序提供初始化值
   P := person{"Tom", 25}
   // 2.通过field:value的方式初始化，这样可以任意顺序
   P := person{age:24, name:"Tom"}
   // 3.new方式,未设置初始值的，会赋予类型的默认初始值
   p := new(person)
   p.age=24
   ```

2. **结构体访问**

   通过`.`操作符用于访问结构的各个字段。

3. **结构体指针**

   指针指向一个结构体
   可以创建**指向结构体的指针**。

   ```go
   var struct_pointer *Books
   ```

   以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前

   ```go
   struct_pointer = &Book1;
   ```

   使用**结构体指针**访问结构体成员，使用==`.` 操作符==

   ```go
   struct_pointer.title;
   ```





# 2.字段

## 2.1匿名字段

1. 在struct中，使用**==不写字段名==**的方式（==只写类型==），使用另一个类型；

   这些字段被称为**匿名字段**。

2. ==其中可以将**匿名字段**理解为**字段名**和**字段类型**都是同一个==

3. 实际就是==字段的继承==

4. 可以使用==`.`操作符==进行调用**匿名字段中的属性值**

5. 若存在匿名字段中的字段与非匿名字段名字相同，则**最外层的优先访问，就近原则**

6. 例子

   ```go
   type Human struct {
       name string
       age int
       weight int
   } 
   type Student struct {
       Human // 匿名字段，那么默认Student就包含了Human的所有字段
       speciality string
   } 
   func main() {
       // 我们初始化一个学生
       mark := Student{Human{"Mark", 25, 120}, "Computer Science"}
       // 我们访问相应的字段
       fmt.Println("His name is ", mark.name)
       fmt.Println("His age is ", mark.age)
       fmt.Println("His weight is ", mark.weight)
       fmt.Println("His speciality is ", mark.speciality)
       // 修改对应的备注信息
       mark.speciality = "AI"
       fmt.Println("Mark changed his speciality")
       fmt.Println("His speciality is ", mark.speciality)
       // 修改他的年龄信息
       fmt.Println("Mark become old")
       mark.age = 46
       fmt.Println("His age is", mark.age)
       // 修改他的体重信息
       fmt.Println("Mark is not an athlet anymore")
       mark.weight += 60
       fmt.Println("His weight is", mark.weight)
   }
   ```

7. 基于上面的理解，所以可以`mark.Human = Human{"Marcus", 55, 220}`和`mark.Human.age -= 1`

8. ==不仅仅是struct字段哦，所有的内置类型和自定义类型都是可以作为匿名字段的。==



## 2.2struct嵌套

1. struct支持嵌套

2. 例子

   ```go
   type Address struct {  
       city, state string
   }
   type Person struct {  
       name string
       age int
       address Address
   }
   
   func main() {  
       var p Person
       p.name = "Naveen"
       p.age = 50
       p.address = Address {
           city: "Chicago",
           state: "Illinois",
       }
       fmt.Println("Name:", p.name)
       fmt.Println("Age:",p.age)
       fmt.Println("City:",p.address.city)
       fmt.Println("State:",p.address.state)
   }
   ```

   

## 2.3提升字段

1. 在结构体中**属于==匿名结构体==的字段**称为**提升字段，**因为它们可以被==直接==访问，就好像它们属于拥有匿名结构字段的结构一样。

2. 例子

   ```go
   // 和上个例子对比
   type Address struct {  
       city, state string
   }
   type Person struct {  
       name string
       age  int
       Address
   }
   
   func main() {  
       var p Person
       p.name = "Naveen"
       p.age = 50
       p.Address = Address{
           city:  "Chicago",
           state: "Illinois",
       }
       fmt.Println("Name:", p.name)
       fmt.Println("Age:", p.age)
       fmt.Println("City:", p.city) //city is promoted field
       fmt.Println("State:", p.state) //state is promoted field
   }
   ```



## 2.4导出struct和字段

1. 如果**结构体**类型以**大写字母开头**，那么它是一个==导出类型==，可以==从其他包访问它==。类似地，如果**结构体的字段**以**大写开头**，则可以从其他包访问它们。

2. 示例代码：

   1.在computer目录下，创建文件spec.go

   ```go
   package computer
   
   type Spec struct { //exported struct  
       Maker string //exported field
       model string //unexported field
       Price int //exported field
   }
   ```

   2.创建main.go 文件

   ```go
   package main
   
   import "structs/computer"  
   import "fmt"
   
   func main() {  
       var spec computer.Spec
       spec.Maker = "apple"
       spec.Price = 50000
       fmt.Println("Spec:", spec)
   }
   ```

   > 目录结构如下：
   >
   > ```shell
   > src  
   >     structs
   >         computer
   >             spec.go
   >         main.go
   > ```



## 2.5struct比较

1. 结构体是**值类型**，如果**每个字段具有可比性**，则是可比较的。如果它们对应的字段相等，则认为两个结构体变量是相等的。
2. **如果结构变量包含的字段是不可比较的，那么结构变量是不可比较的**







# 3.make和new

1. `make`用于**内建类型（map、slice 和channel）的内存分配**。

   `new`用于**各种类型的内存分配**。

2. 内建函数`new`本质上说跟其它语言中的同名函数功能一样：new(T)分配了**零值填充**的T类型的内存空间，并且返回其地址，即一个`*T`类型的值。用Go的术语说，它==返回了一个指针，指向新分配的类型T的零值==**。有一点非常重要：new返回**指针

3. 内建函数`make(T, args)`与`new(T)`有着不同的功能，**make只能创建slice、map和channel**，并且**返回**一个有==初始值(非零)的T类型==，而不是*T。本质来讲，导致这三个类型有所不同的**原因**是==指向数据结构的引用在使用前必须被初始化==。

   例如，一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil。对于slice、map和channel来说，make初始化了内部的数据结构，填充适当的值。

   make返回初始化后的**（非零）值**。