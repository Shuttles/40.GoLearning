package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	// 获取问候信息并打印出来
	message := greetings.Hello("chenzheyu")
	fmt.Println(message)
}
