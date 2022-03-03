# Channel intro

1. Go 对并发的原生支持可不是仅仅停留在口号上的，Go 在语法层面将并发原语 channel 作为**一等公民**对待。
2. 这意味着我们可以像**使用普通变量**那样**使用 channel**，比如，定义 channel 类型变量、给 channel 变量赋值、将 channel 作为参数传递给函数 / 方法、将 channel 作为返回值从函数 / 方法中返回，甚至将 channel 发送到其他 channel 中。这就大大简化了 channel 原语的使用，提升了我们开发者在做并发设计和实现时的体验。





关闭Channel

1. 发送端负责关闭channel，因为发送端没有像接受端那样的、可以安全判断 channel 是否被关闭了的方法。



# Select

1. 当涉及**同时对多个 channel** 进行操作时，我们会结合 Go 为 CSP 并发模型提供的另外一个原语 **select**，一起使用。

2. 通过 select，我们可以同时在**多个** channel 上进行发送 / 接收操作

   ```go
   select {
   case x := <-ch1:     // 从channel ch1接收数据
     ... ...
   
   case y, ok := <-ch2: // 从channel ch2接收数据，并根据ok值判断ch2是否已经关闭
     ... ...
   
   case ch3 <- z:       // 将z值发送到channel ch3中:
     ... ...
   
   default:             // 当上面case中的channel通信均无法实施时，执行该默认分支
   }
   ```

   

3. 当 select 语句中没有 default 分支，而且所有 case 中的 channel 操作都阻塞了的时候，整个 select 语句都将被阻塞，直到某一个 case 上的 channel 变成可发送，或者某个 case 上的 channel 变成可接收，select 语句才可以继续进行下去。





# 无缓冲chan惯用法

无缓冲 channel 兼具==**通信**==和==**同步**==特性，在并发程序中应用颇为广泛。现在我们来看看几个无缓冲 channel 的典型应用：

## 1.用作信号传递

1. 无缓冲 channel 用作信号传递的时候，有两种情况，分别是 1 对 1 通知信号和 1 对 n 通知信号。

2. 1对1

   ```go
   type signal struct{}
   
   func worker() {
       println("worker is working...")
       time.Sleep(1 * time.Second)
   }
   
   func spawn(f func()) <-chan signal {
       c := make(chan signal)
       go func() {
           println("worker start to work...")
           f()
           c <- signal{}
       }()
       return c
   }
   
   func main() {
       println("start a worker...")
       c := spawn(worker)
       <-c
       fmt.Println("worker work done!")
   }
   ```

   

3. 有些时候，无缓冲 channel 还被用来实现 1 对 n 的信号通知机制。这样的信号通知机制，常被用于协调多个 Goroutine 一起工作，比如下面的例子：

   我们可以看到，**关闭一个无缓冲 channel** 会<u>让所有阻塞在这个 channel 上的接收操作返回</u>，从而实现了一种 1 对 n 的“**广播”**机制。





## 2.用于代替锁机制

1. 无缓冲 channel 具有**同步**特性，这让它在某些场合可以**替代锁**，让我们的程序更加清晰，可读性也更好。
2. 这种并发设计逻辑更符合 Go 语言所倡导的“==不要通过<u>共享内存来通信</u>，而是<u>通过通信来共享内存</u>==”的原则。





# 带缓冲chan惯用法

1. 带缓冲的 channel 与无缓冲的 channel 的最大不同之处，就在于它的**异步性**。
2. 也就是说，对一个带缓冲 channel，在缓冲区未满的情况下，对它进行发送操作的 Goroutine 不会阻塞挂起；在缓冲区有数据的情况下，对它进行接收操作的 Goroutine 也不会阻塞挂起。
3. 这种特性让带缓冲的 channel 有着与无缓冲 channel 不同的应用场合。



## 1.用作消息队列



## 2.用作计数信号量

也就是counting semaphore





# 与Select结合的惯用法

channel 和 select 的结合使用能形成强大的表达能力，我们在前面的例子中已经或多或少见识过了。这里我再强调几种 channel 与 select 结合的惯用法。



## 1.利用default分支避免阻塞

1. select 语句的 default 分支的语义，就是在其他非 default 分支因通信未就绪，而无法被选择的时候执行的，这就给 default 分支赋予了一种“避免阻塞”的特性。

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



## 2.实现超时机制

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

3. 不过，在应用带有超时机制的 select 时，我们要特别注意 timer 使用后的释放，尤其在大量创建 timer 的时候。



## 3.实现心跳机制

1. 结合 time 包的 Ticker，我们可以实现带有心跳机制的 select。这种机制让我们可以在监听 channel 的同时，执行一些周期性的任务，比如下面这段代码：

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

2. 这里我们使用 time.NewTicker，创建了一个 Ticker 类型实例 heartbeat。这个实例包含一个 channel 类型的字段 C，这个字段会按一定时间间隔持续产生事件，就像“心跳”一样。这样 for 循环在 channel c 无数据接收时，会每隔特定时间完成一次迭代，然后回到 for 循环进行下一次迭代。

3. 和 timer 一样，我们在使用完 ticker 之后，也不要忘记调用它的 Stop 方法，避免心跳事件在 ticker 的 channel（上面示例中的 heartbeat.C）中持续产生。





# 总结

1. 最后，select 是 Go 为了支持同时操作多个 channel，而引入的另外一个并发原语，select 与 channel 有几种常用的固定搭配，你也要好好掌握和理解。