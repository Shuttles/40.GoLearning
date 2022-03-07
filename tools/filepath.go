package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	var s1 string = "/secrets"
	var s2 string = "ssh-publickey"
	publicKeyFile := filepath.Join(s1, s2)
	fmt.Println(publicKeyFile)
}
