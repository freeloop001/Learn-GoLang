// go run 2.go 18 20
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
