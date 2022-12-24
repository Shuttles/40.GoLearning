# Go线程池



# 为什么需要Goroutine池？

1. 我们就说过：相对于操作系统线程，**Goroutine 的开销十分小**，一个 Goroutine 的起始栈大小为 2KB，而且创建、切换与销毁的代价很低，我们可以创建成千上万甚至更多 Goroutine。

   所以和其他语言不同的是，Go 应用通常可以为每个新建立的连接创建一个对应的新 Goroutine，<u>甚至是为每个传入的请求生成一个 Goroutine 去处理</u>。

   不过，Goroutine 的开销虽然“廉价”，但也不是免费的。最明显的，**一旦规模化后，这种非零成本也会成为瓶颈**。我们以一个 Goroutine 分配 2KB 执行栈为例，100w Goroutine 就是 2GB 的内存消耗。

2. 其次，Goroutine 从Go 1.4 版本开始采用了**连续栈**的方案，但连续栈的原理也决定了，一旦 Goroutine 的执行栈发生了 grow，那么即便这个 Goroutine 不再需要那么大的栈空间，这个 Goroutine 的栈空间也**不会被 Shrink（收缩）**了，这些空间可能会处于长时间闲置的状态，直到 Goroutine 退出。（==不懂这和goroutine池有什么关系==）

3. 另外，随着 Goroutine 数量的增加，<u>Go 运行时进行 Goroutine 调度的处理器消耗，也会随之增加</u>，成为阻碍 Go 应用性能提升的重要因素。

4. 那么面对这样的问题，常见的应对方式是什么呢？

   Goroutine 池就是一种常见的解决方案。这个方案的**核心思想**是==对 Goroutine 的重用==，也就是<u>把 M 个计算任务调度到 N 个 Goroutine 上，而不是为每个计算任务分配一个独享的 Goroutine</u>，从而提高计算资源的利用率。





# 1.Pool结构定义

workpool的实现主要分为三个部分：

1. pool的创建与销毁
2. pool中worker（goroutine）的管理
3. task的提交与调度

![](https://static001.geekbang.org/resource/image/d4/fd/d48ba3a204ca6e8961a4425573afa0fd.jpg?wh=1920x1047)

```go
type Pool struct {
  capacity int

  active chan struct{}
  tasks chan Task
  
  wg sync.WaitGroup
  quit chan struct{} // 用于通知各个worker退出
}
```









## 1.1对worker的管理

1. capacity 是 pool 的一个属性，代表整个 pool 中 worker 的最大容量。

2. 我们使用一个带缓冲的 channel：**active**，作为 worker 的“计数器”

   当 active channel 可写时，我们就创建一个 worker，用于处理用户通过 Schedule 函数提交的待处理的请求。

   当 active channel 满了的时候，pool 就会停止 worker 的创建，直到某个 worker 因故退出，active channel 又空出一个位置时，pool 才会创建新的 worker 填补那个空位。



## 1.2task的提交与调度

1. 我们把用户要提交给 workerpool 执行的请求抽象为一个 Task。

2. Task 的提交与调度也很简单：Task 通过 **Schedule** 函数提交到一个 task channel 中，<u>已经创建的 worker 将从这个 task channel 中读取 task 并执行</u>。

3. Task 是一个对用户提交的请求的抽象，它的本质就是一个函数类型：

   ```go
   type Task func()
   ```

   这样，用户通过 Schedule 方法实际上提交的是一个函数类型的实例。







workerpool包对外主要提供三个API，它们分别是：

1. `workerpool.New`:用于**创建**一个Pool类型实例，并将Pool池的worker管理机制运行起来（**run方法**）

   + New()，接收一个参数**capacity**用于制定workerpool池的容量，控制workerpool最多只能有capacity个worker，共同处理用户提交的task

   + Pool类型实例p完成初始化后，我们创建一个新的goroutine，用于对workerpool进行管理，这个goroutine执行的是Pool类型的**run方法**

   + run()方法内是一个无限循环，循环体中使用**select**监视Pool实例中的两个channel：quit和acitve。（这种在for中用select监视多个channel的实现，在Go代码中十分常见）

     <u>当接收到来自quit的退出信号时，这个goroutine就会结束运行；而当active channel可写时，run方法就会创建出一个新的worker goroutine</u>。（为了方便在程序中区分各个 worker 输出的日志，我这里将一个从 1 开始的变量 idx 作为 worker 的编号，并把它以参数的形式传给创建 worker 的方法。）

   + 我们再将创建新的worker gorouine的职责，封装到名为newWorker的方法中。

     由于每个 worker 运行于一个独立的 Goroutine 中，newWorker 方法通过 go 关键字创建了一个新的 Goroutine 作为 worker。

     新 worker 的核心，依然是一个基于 **for-select** 模式的循环语句，在循环体中，新 worker 通过 **select** 监视 quit 和 tasks 两个 channel。和前面的 run 方法一样，当接收到来自 quit channel 的退出“信号”时，这个 worker 就会结束运行。

     tasks channel 中放置的是用户通过 **Schedule** 方法提交的请求，新 worker 会从这个 channel 中获取最新的 Task 并运行这个 Task。

     在新 worker 中，为了防止用户提交的 task 抛出 panic，进而导致整个 workerpool 受到影响，我们在 worker 代码的开始处，使用了 defer+recover 对 panic 进行捕捉，捕捉后 worker 也是要退出的，于是我们还通过<-p.active更新了 worker 计数器。并且一旦 worker goroutine 退出，p.wg.Done 也需要被调用，这样可以减少 WaitGroup 的 Goroutine 等待数量。

2. `workerpool.Free`:用于**销毁**一个Pool池，停掉所有Pool池中的worker

3. `Pool.Schedule`:这是Pool类型的一个导出方法，workerpool包的用户<u>通过该方法向Pool池提交代执行的任务（Task）</u>

   + Schedule 方法的核心逻辑，是将传入的 Task 实例发送到 workerpool 的 tasks channel 中。但考虑到现在 workerpool 已经被销毁的状态，我们这里通过一个 select，检视 quit channel 是否有“信号”可读，如果有，就返回一个哨兵错误 ErrWorkerPoolFreed。如果没有，一旦 p.tasks 可写，提交的 Task 就会被写入 tasks channel，以供 pool 中的 worker 处理。
   + 这里要注意的是，这里的 Pool 结构体中的 tasks 是一个无缓冲的 channel，如果 pool 中 worker 数量已达上限，而且 worker 都在处理 task 的状态，那么 Schedule 方法就会阻塞，直到有 worker 变为 idle 状态来读取 tasks channel，schedule 的调用阻塞才会解除。







# Main.go

```go
package main
  
import (
    "time"
    "github.com/bigwhite/workerpool"
)

func main() {
    p := workerpool.New(5)

    for i := 0; i < 10; i++ {
        err := p.Schedule(func() {
            time.Sleep(time.Second * 3)
        })
        if err != nil {
            println("task: ", i, "err:", err)
        }
    }

    p.Free()
}
```

这个示例程序创建了一个 capacity 为 5 的 workerpool 实例，并连续向这个 workerpool 提交了 10 个 task，每个 task 的逻辑很简单，**只是 Sleep 3 秒后就退出**。

main 函数在提交完任务后，调用 workerpool 的 Free 方法销毁 pool，pool 会等待所有 worker 执行完 task 后再退出。



demo1示例的运行结果如下：

```shell

workerpool start
worker[005]: start
worker[005]: receive a task
worker[003]: start
worker[003]: receive a task
worker[004]: start
worker[004]: receive a task
worker[001]: start
worker[002]: start
worker[001]: receive a task
worker[002]: receive a task
worker[004]: receive a task
worker[005]: receive a task
worker[003]: receive a task
worker[002]: receive a task
worker[001]: receive a task
worker[001]: exit
worker[005]: exit
worker[002]: exit
worker[003]: exit
worker[004]: exit
workerpool freed
```

