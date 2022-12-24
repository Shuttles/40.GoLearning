# Channel

原文地址：https://time.geekbang.org/column/article/477365



# 1.前言

1. channel既可以用来实现 Goroutine 间的==通信==，还可以实现 Goroutine 间的==同步==。



# 2.channel是一等公民

1. 这意味着我们可以像使用普通变量那样使用 channel，比如，

   + 定义 channel 类型变量、

   + 给 channel 变量赋值、

   + 将 channel 作为参数传递给函数 / 方法、

   + 将 channel 作为返回值从函数 / 方法中返回，

   + 甚至将 channel 发送到其他 channel 中。

     这就大大简化了 channel 原语的使用，提升了我们开发者在做并发设计和实现时的体验。



## 2.1创建

1. 和切片、结构体、map 等一样，channel 也是一种复合数据类型。也就是说，我们在声明一个 channel 类型变量时，必须给出其具体的元素类型，比如

   ```go
   var ch chan int
   // 声明了一个元素为 int 类型的 channel 类型变量 ch
   ```

2. 如果 channel 类型变量在声明时没有被赋予初值，那么它的默认值为 nil。

3. 和其他复合数据类型支持使用复合类型字面值作为变量初始值不同，为 channel 类型变量**赋初值的唯一方法**就是使用 ==make== 这个 Go 预定义的函数，比如下面代码：

   ```go
   ch1 := make(chan int)
   ch2 := make(chan int, 5)
   ```

   第一行我们通过make(chan T)创建的、元素类型为 T 的 channel 类型，是**无缓冲 channel**，

   而第二行中通过带有 capacity 参数的make(chan T, capacity)创建的元素类型为 T、缓冲区长度为 capacity 的 channel 类型，是**带缓冲 channel**。

4. 这两种类型的变量关于发送（send）与接收（receive）的特性是不同的，



## 2.2发送与接收

1. Go 提供了<-操作符用于对 channel 类型变量进行发送与接收操作：

   ```go
   ch1 <- 13    // 将整型字面值13发送到无缓冲channel类型变量ch1中
   n := <- ch1  // 从无缓冲channel类型变量ch1中接收一个整型值存储到整型变量n中
   ch2 <- 17    // 将整型字面值17发送到带缓冲channel类型变量ch2中
   m := <- ch2  // 从带缓冲channel类型变量ch2中接收一个整型值存储到整型变量m中
   ```

2. 在理解 channel 的发送与接收操作时，你一定要始终牢记：channel 是用于 Goroutine 间通信的，所以<u>绝大多数对 channel 的读写都被分别放在了==不同的 Goroutine 中==</u>。

3. 由于无缓冲 channel 的运行时层实现**不带有缓冲区**，所以 Goroutine 对无缓冲 channel 的接收和发送操作是同步的。也就是说，对同一个无缓冲 channel，只有对它进行接收操作的 Goroutine 和对它进行发送操作的 Goroutine ==都存在==的情况下，通信才能得以进行，否则单方面的操作会让对应的 Goroutine <u>陷入挂起状态</u>

4. 对**无缓冲** channel 类型的发送与接收操作，一定要放在**两个不同**的 Goroutine 中进行，否则会导致 deadlock。

5. 和无缓冲 channel 相反，**带缓冲** channel 的运行时层实现带有缓冲区，因此，对带缓冲 channel 的发送操作在缓冲区未满、接收操作在缓冲区非空的情况下是**异步**的（发送或接收**不需要阻塞等待**）。



## 2.3Only类型Channel

1. 使用操作符<-，我们还可以声明**只发送(==只向它发送==)** channel 类型（send-only）和**只接收（==只从它接收==）** channel 类型（recv-only）

   ```go
   ch1 := make(chan<- int, 1) // 只发送channel类型
   ch2 := make(<-chan int, 1) // 只接收channel类型
   
   <-ch1       // invalid operation: <-ch1 (receive from send-only type chan<- int)
   ch2 <- 13   // invalid operation: ch2 <- 13 (send to receive-only type <-chan int)
   ```



## 2.4关闭Channel

1. channel 关闭后，所有等待从这个 channel 接收数据的操作都将返回。

2. 采用不同接收语法形式的语句，在 channel 被关闭后的返回值的情况：

   ```go
   n := <- ch      // 当ch被关闭后，n将被赋值为ch元素类型的零值
   m, ok := <-ch   // 当ch被关闭后，m将被赋值为ch元素类型的零值, ok值为false
   for v := range ch { // 当ch被关闭后，for range循环结束
       ... ...
   }
   ```

3. 通过“comma, ok”惯用法或 for range 语句，我们可以准确地判定 channel 是否被关闭。

   而单纯采用n := <-ch形式的语句，我们就无法判定从 ch 返回的元素类型零值，究竟是不是因为 channel 被关闭后才返回的。

4. channel 的一个使用惯例，就是**发送端负责关闭 channel**，因为发送端没有像接受端那样的、可以安全判断 channel 是否被关闭了的方法。



## 2.5Select

1. 当涉及同时对多个 channel 进行操作时，我们会结合 Go 为 CSP 并发模型提供的另外一个原语 select，一起使用。

2. 通过 select，我们可以同时在多个 channel 上进行发送 / 接收操作：

   ```go
   select {
   case x := <-ch1;   // 从channel ch1接收数据 
     ... ...
   case y, ok := <-ch2;    // 从channel ch2接收数据，并根据ok值判断ch2是否已经关闭 
     ... ...
   case ch3 <- z;     // 将z值发送到channel ch3中: 
     ... ...
   default: // 当上面case中的channel通信均无法实施时，执行该默认分支
   }
   ```

3. 当 select 语句中没有 default 分支，而且所有 case 中的 channel 操作都阻塞了的时候，整个 select 语句都将**被阻塞**，直到某一个 case 上的 channel 变成可发送，或者某个 case 上的 channel 变成可接收，select 语句才可以继续进行下去。

4. channel 和 select 两种原语的操作都十分简单，它们都遵循了 Go 语言“==追求简单==”的设计哲学，但它们却为 Go 并发程序带来了强大的表达能力。





# 3.无缓冲channel的惯用法

无缓冲 channel 兼具==通信==和==同步==特性，在并发程序中应用颇为广泛。现在我们来看看几个无缓冲 channel 的典型应用：

## 3.1用作信号传递

1. 有两种情况，分别是 **1 对 1 通知信号**和 **1 对 n 通知信号**。我们先来分析下 1 对 1 通知信号这种情况。

   例子在原文链接中

   ```go
   func main() { 
     println("start a worker...") 
     c := spawn(worker) <-c 
     fmt.Println("worker work done!")
   }
   ```

   

   在这个例子中，spawn 函数返回的 channel，被用于承载<u>新 Goroutine 退出的“通知信号”</u>，这个信号专门用作通知 main goroutine。main goroutine 在调用 spawn 函数后一直阻塞在对这个“通知信号”的接收动作上。

2. 有些时候，无缓冲 channel 还被用来实现 **1 对 n** 的信号通知机制。这样的信号通知机制，常被用于<u>协调多个 Goroutine</u> 一起工作

   ```go
   func main() {
       fmt.Println("start a group of workers...")
       groupSignal := make(chan signal)
       c := spawnGroup(worker, 5, groupSignal)
       time.Sleep(5 * time.Second)
       fmt.Println("the group of workers start to work...")
       close(groupSignal)
       <-c
       fmt.Println("the group of workers work done!")
   }
   ```

   这个例子中，main goroutine 创建了一组 5 个 worker goroutine，这些 Goroutine 启动后会阻塞在名为 groupSignal 的无缓冲 channel 上。

   main goroutine 通过close(groupSignal)<u>向所有 worker goroutine 广播“开始工作”的信号</u>，收到 groupSignal 后，所有 worker goroutine 会“同时”开始工作，就像起跑线上的运动员听到了裁判员发出的起跑信号枪声。

   我们可以看到，**关闭一个无缓冲 channel 会让所有阻塞在这个 channel 上的接收操作返回**，从而实现了一种 1 对 n 的“==广播==”机制。



## 3.2用于替代锁机制

1. 无缓冲 channel 具有**同步**特性，这让它在某些场合可以替代锁，让我们的程序更加清晰，可读性也更好。

2. 首先我们看一个传统的、基于“共享内存”+“互斥锁”的 Goroutine 安全的计数器的实现：

   /Users/chenbin/40.GoLearning/5.goroutine/unBufferChan/counter1.go中

   在这个示例中，我们使用了一个<u>带有互斥锁保护的全局变量作为计数器</u>，所有要操作计数器的 Goroutine 共享这个全局变量，并<u>在互斥锁的同步下对计数器进行自增操作</u>。

3. 接下来我们再看更符合 Go 设计惯例的实现，也就是**使用无缓冲 channel 替代锁**后的实现：

   /Users/chenbin/40.GoLearning/5.goroutine/unBufferChan/counter2.go中

   在这个实现中，我们将计数器操作全部交给一个独立的 Goroutine 去处理，并<u>通过无缓冲 channel 的同步阻塞特性，实现了计数器的控制</u>。

   这样其他 Goroutine 通过 Increase 函数试图增加计数器值的动作，实质上就<u>转化为了一次无缓冲 channel 的接收动作</u>。

   这种并发设计逻辑更符合 Go 语言所倡导的“==不要通过共享内存来通信，而是通过通信来共享内存==”的原则。

   

   



# 4.带缓冲channel的惯用法、

带缓冲的 channel 与无缓冲的 channel 的最大不同之处，就在于它的==异步性==。也就是说，对一个带缓冲 channel，在缓冲区未满的情况下，对它进行发送操作的 Goroutine 不会阻塞挂起；在缓冲区有数据的情况下，对它进行接收操作的 Goroutine 也不会阻塞挂起。



## 4.1用作消息队列

要注意的是，Go 支持 channel 的初衷是将它作为 <u>Goroutine 间的通信手段</u>，它并不是专门用于消息队列场景的。如果你的项目需要专业消息队列的功能特性，比如支持优先级、支持权重、支持离线持久化等，那么 channel 就不合适了，可以使用第三方的专业的消息队列实现。



## 4.2用作计数信号量

1. Go 并发设计的一个惯用法，就是将带缓冲 channel 用作计数信号量（counting semaphore）。
2. 带缓冲 channel 中的当前数据个数代表的是，**当前同时处于活动状态（处理业务）的 Goroutine 的数量**，而带缓冲 channel 的**容量**（capacity），就代表了允许同时处于活动状态的 Goroutine 的**最大数量**。
3. 向带缓冲 channel 的一个<u>发送操作表示获取一个信号量</u>，而从 channel 的一个<u>接收操作则表示释放一个信号量。</u>







# 5.len(channel)

1. **len** 是 Go 语言的一个内置函数，它支持接收数组、切片、map、字符串和 channel 类型的参数，并返回对应类型的“长度”，也就是一个整型值。

2. 针对 channel ch 的类型不同，len(ch) 有如下两种语义：

   + 当 ch 为无缓冲 channel 时，len(ch) 总是返回 0；
   + 当 ch 为带缓冲 channel 时，len(ch) 返回当前 channel ch 中尚<u>未被读取的元素个数</u>。

3. channel 原语用于多个 Goroutine 间的通信，<u>一旦多个 Goroutine 共同对 channel 进行收发操作</u>，len(channel) 就会在多个 Goroutine 间形成“**竞态**”。

   <u>单纯地依靠 len(channel) 来判断 channel 中元素状态</u>，是**不能**保证在后续对 channel 的收发时 channel 状态是不变的。

4. 有些时候我们想在不改变 channel 状态的前提下，**单纯地侦测 channel 的状态**，而又不会因 channel 满或空阻塞在 channel 上。但很遗憾，目前**没有**一种方法可以在实现这样的功能的同时，适用于所有场合。

5. 但是在特定的场景下，我们可以用 len(channel) 来实现。比如下面这两种场景：

   ![](https://static001.geekbang.org/resource/image/b3/37/b31d081fcced758b8f99c938a0b75237.jpg?wh=1920x1047)





# 6.nil channel

1. 如果一个 channel 类型变量的值为 nil，我们称它为 nil channel。nil channel 有一个特性，那就是对 nil channel 的读写都会发生阻塞。比如下面示例代码：

   ```go
   func main() { 
     var c chan int <-c //阻塞
   }
              
   func main() { 
     var c chan int c<-1 //阻塞
   }
   ```

2. 不过，nil channel 的这个特性可不是一无是处，有些时候应用 nil channel 的这个特性可以得到事半功倍的效果。

   我们知道，<u>对一个 nil channel 执行获取操作，这个操作将阻塞</u>。所以如果我们想阻塞，那就==将chan赋值为nil==





# 7.Select

1. channel 和 select 的结合使用能形成强大的表达能力，我们在前面的例子中已经或多或少见识过了。

   这里我再强调几种 channel 与 select 结合的惯用法。



## 7.1利用default分支避免阻塞

1. select 语句的 default 分支的语义，就是在其他非 default 分支因通信未就绪，而无法被选择的时候执行的，这就给 default 分支赋予了一种“**避免阻塞**”的特性。

2. 其实在前面的“len(channel) 的应用”小节的例子中，我们就已经用到了“利用 default 分支”实现的trySend和tryRecv两个函数：

   ```go
   func tryRecv(c <-chan int) (int, bool) {
     select {
     case i := <-c:
       return i, true
   
     default: // channel为空
       return 0, false
     }
   }
   
   func trySend(c chan<- int, i int) bool {
     select {
     case c <- i:
       return true
     default: // channel满了
       return false
     }
   }
   ```

3. 而且，无论是无缓冲 channel 还是带缓冲 channel，这两个函数都能适用，并且不会阻塞在空 channel 或元素个数已经达到容量的 channel 上。



## 7.2实现超时机制

1. 带超时机制的 select，是 Go 中常见的一种 select 和 channel 的组合用法。通过超时事件，我们既可以避免长期陷入某种操作的等待中，也可以做一些异常处理工作。

2. 比如，下面示例代码实现了一次具有 30s 超时的 select：

   ```go
   func worker() {
     select {
     case <-c:
          // ... do some stuff
     case <-time.After(30 *time.Second):
         return
     }
   }
   ```

   



## 7.3实现心跳机制

1. 结合 time 包的 Ticker，我们可以实现带有心跳机制的 select。这种机制<u>让我们可以在监听 channel 的同时，执行一些周期性的任务，</u>

2. 比如下面这段代码：

   ```go
   func worker() {
     heartbeat := time.NewTicker(30 * time.Second)
     defer heartbeat.Stop()
     for {
       select {
       case <-c:
         // ... do some stuff
       case <- heartbeat.C:
         //... do heartbeat stuff
       }
     }
   }
   ```

3. 这里我们使用 time.NewTicker，创建了一个 Ticker 类型实例 heartbeat。

   这个实例包含一个 channel 类型的字段 C，这个字段会按一定时间间隔持续产生事件，就像“**心跳**”一样。

   这样 for 循环在 channel c 无数据接收时，会每隔特定时间完成一次迭代，然后回到 for 循环进行下一次迭代。

   和 timer 一样，我们在使用完 ticker 之后，也不要忘记调用它的 Stop 方法，避免心跳事件在 ticker 的 channel（上面示例中的 heartbeat.C）中持续产生。



# 总结

1. Go为了原生支持并发，把channel视作**一等公民**身份，这就大幅提升了开发人员使用channel进行并发设计和实现的体验

2. 通过**预定义函数 make**，我们可以创建两类 channel：无缓冲 channel 与带缓冲的 channel。

3. 这两类 channel 具有不同的收发特性，可以适用于不同的应用场合：

   + **无缓冲** channel 兼具**通信与同步**特性，常用于作为<u>信号通知或替代同步锁</u>；
   + 而**带缓冲** channel 的**异步性**，让它更适合用来实现基于内存的<u>消息队列、计数信号量</u>等。

4. 要牢记值为 nil 的 channel 的**阻塞**特性，有些时候它也能帮上大忙。

   而面对已关闭的 channel 你也一定要小心，尤其要<u>避免向已关闭的 channel 发送数据，那会导致 panic</u>。

5. select 是 Go 为了支持同时操作多个 channel，而引入的另外一个并发原语，select 与 channel 有几种**常用的固定搭配**，要好好掌握和理解。