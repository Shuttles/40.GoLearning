package main

import (
	"fmt"
	"strings"
)

func main() {
	ans := "rsync-wsts-dest-0.default.pod.transwarp.local"
	hostname := strings.Join([]string{"Hostname", "Namespace", "pod.transwarp.local"}, ".")
	fmt.Println(hostname)
	fmt.Println(ans)
}
