package main

import (
	"time"
	"sync"
	"log"
)

var active = make(chan struct{}, 3)
var jobs = make(chan int, 10)


func main() {
	go func() {
		for i := 0; i < 8; i++ {
			jobs <- (i + 1)
		}
		// 为什么下面一行不会导致下面的goroutine错误？
		// 可以推测，带缓冲channel，close掉之后还可以读取原来的内容
		// 如果没有下面这行，会导致死锁
		close(jobs) 
	}()

	var wg sync.WaitGroup

	for j := range jobs {
		wg.Add(1)
		go func(j int) {
			active <- struct{}{}
			log.Printf("handle jobs: %d\n", j)
			time.Sleep(2 * time.Second)
			<-active
			wg.Done()
		}(j)
	}
	wg.Wait()
}