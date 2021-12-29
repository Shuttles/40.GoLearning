**通道**可以被认为是==Goroutines通信的管道==。类似于管道中的水从一端到另一端的流动，数据可以从一端发送到另一端，通过通道接收。

# 1.声明Channel

1. **每个通道都有与其相关的类型**。该类型是**通道允许传输的数据类型**。

2. 通道的零值为nil。**nil通道没有任何用处**，因此通道**必须**使用类似于map和切片的方法来定义（==make==）。

3. 示例代码：

   ```go
   package main
   
   import "fmt"
   
   func main() {  
       var a chan int
       if a == nil {
           fmt.Println("channel a is nil, going to define it")
           a = make(chan int)
           fmt.Printf("Type of a is %T", a)
       }
   }
   ```

   运行结果：

   ```
   channel a is nil, going to define it  
   Type of a is chan int
   ```

   **也可以简短的声明：**

   ```go
   a := make(chan int)
   ```



# 2.发送和接受

1. 发送和接收的语法：

   ```go
   data := <- a // read from channel a  
   a <- data // write to channel a
   ```

   在通道上**箭头的方向指定数据是发送还是接收**。

2. 发送和接受默认是**阻塞的**

   当一个数据被发送到通道时，在发送语句中被阻塞，直到另一个Goroutine从该通道读取数据。类似地，当从通道读取数据时，读取被阻塞，直到一个Goroutine将数据写入该通道。

   这些通道的特性是帮助Goroutines**有效地进行通信**，而无需像使用其他编程语言中非常常见的显式锁或条件变量。

3. 示例代码：

   ```go
   package main
   
   import (  
       "fmt"
   )
   
   func hello(done chan bool) {  
       fmt.Println("Hello world goroutine")
       done <- true
   }
   func main() {  
       done := make(chan bool)
       go hello(done)
       <-done // 接收数据，阻塞式
       fmt.Println("main function")
   }
   ```

   运行结果：

   ```
   Hello world goroutine  
   main function
   ```

   在上面的程序中，我们在第一行中创建了一个done bool通道。把它作为参数传递给hello Goroutine。第14行我们正在接收已完成频道的数据。这一行代码是**阻塞**的，这意味着在某些Goroutine将数据写入到已完成的通道之前，程序将不会执行到下一行代码。因此，这就消除了对时间的需求。

   > 代码`<-done`接收来自done Goroutine的数据，但不使用或存储任何变量中的数据。这是**完全合法**的。

   现在，我们的main Goroutine阻塞等待已完成通道的数据。hello Goroutine接收这个通道作为参数，打印hello world Goroutine，然后写入done通道。当此写入完成时，main的Goroutine接收来自已完成通道的数据，然后输出文本主函数。



# 3.死锁

1. 使用通道时要考虑的一个重要因素是死锁。



# 4.定向通道

1. 之前我们学习的通道都是双向通道，我们可以通过这些通道接收或者发送数据。我们也可以创建单向通道，这些通道只能发送或者接收数据。

   创建仅能发送数据的通道，示例代码：

   ```go
   package main
   
   import "fmt"
   
   func sendData(sendch chan<- int) {  
       sendch <- 10
   }
   
   func main() {  
       sendch := make(chan<- int)
       go sendData(sendch)
       fmt.Println(<-sendch)
   }
   ```



# 5.关闭chan

1. **发送者**有可以通过**关闭信道**，来通知接收方不会有更多的数据被发送到信道上。

   **接收者**可以在接收来自通道的数据时使用额外的变量来**检查通道是否已经关闭**。

   语法结构：

   ```go
   v, ok := <- ch
   ```

   在上面的语句中，如果ok的值是true，表示成功的从通道中读取了一个数据value。如果ok是false，这意味着我们正在从一个封闭的通道读取数据。从闭通道读取的值将是通道类型的零值。

   例如，**如果通道是一个int通道，那么从封闭通道接收的值将为0。**

2. 示例代码：

   ```go
   package main
   
   import (  
       "fmt"
   )
   
   func producer(chnl chan int) {  
       for i := 0; i < 10; i++ {
           chnl <- i
       }
       close(chnl)
   }
   func main() {  
       ch := make(chan int)
       go producer(ch)
       for {
           v, ok := <-ch
           if ok == false {
               break
           }
           fmt.Println("Received ", v, ok)
       }
   }
   ```

   运行结果

   ```
   Received  0 true  
   Received  1 true  
   Received  2 true  
   Received  3 true  
   Received  4 true  
   Received  5 true  
   Received  6 true  
   Received  7 true  
   Received  8 true  
   Received  9 true
   ```

   在上面的程序中，producer Goroutine将0到9写入chnl通道，然后关闭通道。

   主函数里有一个无限循环。它**检查通道是否在行号中使用变量ok关闭**。如果ok是假的，则意味着通道关闭，因此循环结束。



# 6.chan上的for-range

1. for循环的`for range`形式可用于**从通道接收值，直到它关闭为止。**

2. 使用range循环，示例代码：

   ```go
   package main
   
   import (  
       "fmt"
   )
   
   func producer(chnl chan int) {  
       for i := 0; i < 10; i++ {
           chnl <- i
       }
       close(chnl)
   }
   func main() {  
       ch := make(chan int)
       go producer(ch)
       for v := range ch {
           fmt.Println("Received ",v)
       }
   }
   ```

   



# 7.缓冲通道

1. 之前学习的所有通道基本上都没有缓冲。**发送和接收到一个未缓冲的通道是阻塞的**。

2. 可以**用缓冲区创建一个通道**。

   发送到一个缓冲通道**只有在缓冲区满时才被阻塞**。类似地，从缓冲通道接收的信息**只有在缓冲区为空时才会被阻塞。**

3. 可以通过**将额外的容量参数传递给`make`函数来创建缓冲通道**，该函数**指定缓冲区的大小。**

   语法：

   ```go
   ch := make(chan type, capacity)
   ```

   上述语法的容量应该大于0，以便通道具有缓冲区。默认情况下，无缓冲通道的容量为0，因此在之前创建通道时省略了容量参数。

4. 示例代码：

   ```go
   package main
   
   import (  
       "fmt"
   )
   
   
   func main() {  
       ch := make(chan string, 2)
       ch <- "naveen"
       ch <- "paul"
       fmt.Println(<- ch)
       fmt.Println(<- ch)
   }
   ```