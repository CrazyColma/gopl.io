// exer2 输出参数索引和值 每行一个
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args[1:] {
		fmt.Println(index, ": ", arg)
	}
}
