# 1.Map介绍

1. map是Go中的内置类型，它将**一个值与一个键关**联起来。可以使用相应的**键**（key）检索**值**（value）。
2. Map 是一种集合，所以我们可以像迭代数组和切片那样**迭代**它。不过，Map 是**无序**的（并且每次打印出来的map都会不一样），我们无法决定它的返回顺序，这是因为 Map 是使用 **hash 表**来实现的，也是==引用类型==
3. 内置的**len函数**同样适用于map，返回**map拥有的key的数量**
4. map的key可以是**所有可比较的类型**，如布尔型、整数型、浮点型、复杂型、字符串型……



# 2.创建map

1. 可以使用内建函数 make 也可以使用 map 关键字来定义 Map:

   ```go
   /* 声明变量，默认 map 是 nil */
   var map_variable map[key_data_type]value_data_type
   
   /* 使用 make 函数，就不是nil map了 */
   map_variable = make(map[key_data_type]value_data_type)
   
   // 或者直接用make声明
   var map_variable = make(map[key_type]value_type)
   
   rating := map[string]float32 {"C":5, "Go":4.5, "Python":4.5, "C++":2 }
   ```

   如果不初始化 map，那么就会创建一个 nil map。**nil map 不能用来存放键值对**；**只能用make再创建一遍**

2. 所以，要么

   + **声明(var)的时候初始化**
   + 声明(var)不初始化，**想初始化之前make一下**
   + 直接用**make声明**



# 3.删除元素

1. `delete(map, key)` 函数用于删除集合的元素, 参数为 map 和其对应的 key。
2. `delete`函数**不返回任何值。**



# 4.ok-idiom

1. 我们可以通过key获取map中对应的value值。语法为：

   ```go
   map[key]
   ```

   但是当key如果不存在的时候，我们会得到该value值类型的**默认零值**，比如string类型得到空字符串，int类型得到0。但是程序**不会报错**。

2. 所以我们可以使用**ok-idiom获取值**，**可知道key/value是否存在**

   ```go
   value, ok := map[key]
   ```



# 5.map长度

1. 使用len函数可以确定map的长度。

   ```go
   len(map)  // 可以得到map的长度
   ```



# 6.map是引用类型

1. 与切片相似，映射是**引用类型**。
2. 当将映射分配给一个新变量时，它们都指向**相同的内部数据结构**。因此，**一个的变化会反映另一个。**
3. map不能使用==操作符进行比较。==只能用来检查map是否为空。否则会报错：invalid operation: map1 == map2 (map can only be comparedto nil)