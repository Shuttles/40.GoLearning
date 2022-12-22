package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	publicKeyFile := filepath.Join("/secrets", "ssh-publickey")
	fmt.Println("publicKeyFile=", publicKeyFile)
	privateKeyFile := filepath.Join("/secrets", "ssh-privatekey")
	fmt.Println("privateKey=", privateKeyFile)
}
