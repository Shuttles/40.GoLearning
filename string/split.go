package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "-a      -k -v -r -P 2"
	separator := " "
	parameters := strings.Split(str, separator)
	if parameters == nil {
		fmt.Println("parameters == nil")
		return
	}
	for _, param := range parameters {
		fmt.Println(param)
	}
}
