package main

import (
	"fmt"
	"os"
)

func main() {
	fromToPathJSON := os.Getenv("SyncFromTo")
	fmt.Println("fromToPathJSON = ", fromToPathJSON)

	http_proxy := os.Getenv("http_proxy")
	fmt.Println("http_proxy = ", http_proxy)
	return
}
