package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	params := []string{}
	params = append(params, "-a")
	params = append(params, " ")
	params = append(params, " ")
	params = append(params, "-v")
	params = append(params, "/home/transwarp/40.GoLearning")
	params = append(params, "root@172.18.120.55:/root/zheyuchen/")

	cmd := exec.Command("rsync", params...)
	log.Printf("cmd detail:%v", cmd.Args)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
