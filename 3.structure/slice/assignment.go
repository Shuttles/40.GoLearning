package main

import "fmt"

func modify(s1 []int) {
	s = s1
}

var s = []int{1, 2, 3}

func main() {
	s1 := []int{7, 8, 9}
	modify(s1)
	fmt.Println(s)
	return
}
