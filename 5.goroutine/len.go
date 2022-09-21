package main

import(
	"fmt"
	"time"
	"sync"
)

func tryRecv(c <-chan int) (int, bool) {
	select {
	case val := <-c:
		return val, true
	default:
		return 0, false
	}
}

func trySend(c chan<- int, i int) bool {
	select {
	case c <- i:
		return true
	default:
		return false
	}
}

func producer(c chan<- int) {
	var val int = 1
	for {
		time.Sleep(2 * time.Second)
		ok := trySend(c, val)
		if ok {
			fmt.Printf("[producer]: send [%d] to channel\n", val)
			val++
			continue
		}
		fmt.Printf("[producer]: try send[%d], but channel is full\n", val)
	}
}

func consumer(c <-chan int) {
	for {
		val, ok := tryRecv(c)
		if !ok {
			fmt.Println("[consumer]: try to recv from channel, but the channel if empty")
			time.Sleep(1 * time.Second)
			continue
		}
		fmt.Printf("[consumer]: recv [%d] from channel\n", val)
		if val > 3 {
			fmt.Println("[consumer]: exit")
			return 
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int, 3)
	wg.Add(2)
	go func() {
		producer(c)
		wg.Done()
	}()

	go func() {
		consumer(c)
		wg.Done()
	}()

	wg.Wait()
}