// echo1 输出命令行参数
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	/**
	:=是"短变量声明" 所以不需要手动指定变量类型
	*/
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
