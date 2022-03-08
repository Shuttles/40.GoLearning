package main

import "fmt"

func main() {
	switch x := 4; x {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fallthrough
	default:
		fmt.Println("I am default!!!")
	}

}
