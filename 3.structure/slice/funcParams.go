package main

import "fmt"

func modify(arr []int) []int {
	arr = append(arr, 5)
	return arr
}

func main() {
	s1 := []int{1, 2, 3, 4}
	s2 := modify(s1)
	fmt.Println(s2)
}
