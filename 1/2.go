// go run 2.go 18 20
// 输出命令行参数的index和value
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, name := range os.Args[1:] {
		fmt.Println(index, name)
	}
}
