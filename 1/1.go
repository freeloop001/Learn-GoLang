// go run 1.go
// 输出命令名称
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
}
