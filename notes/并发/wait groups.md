

# 1.背景

1. 很多情况下，我们需要保障一组goroutine**全部结束返回**。这需要借助sync包的WaitGroup来实现。

   

# 2.定义

1. WatiGroup是sync包中的一个struct类型，官方文档的描述是：一个`waitGroup`对象==**等待一组goroutine结束**==。

   ```go
   // A WaitGroup waits for a collection of goroutines to finish. The main
   // goroutine calls Add to set the number of goroutines to wait for. Then each
   // of the goroutines runs and calls Done when finished. At the same time, Wait
   // can be used to block until all goroutines have finished.
   
   // A WaitGroup must not be copied after first use.
   type WaitGroup struct {
           // Has unexported fields.
   }
   
   func (wg *WaitGroup) Add(delta int)
   func (wg *WaitGroup) Done()
   func (wg *WaitGroup) Wait()
   ```

2. 使用方法：

   1. `main goroutine`通过调用 `wg.Add(delta int)` 设置`worker goroutine`的个数，然后创建`worker goroutine`；（<u>相当于+delta</u>）
   2. `worker goroutine`执行**结束**以后，都要调用 `wg.Done()`  （建议用**defer**）（<u>相当于-1</u>）；
   3. `main gorouine`调用 `wg.Wait()` 且被block，直到所有`worker goroutine`全部执行结束后返回。



# 3.最佳实践

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		i := i // interesting

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
}
```

