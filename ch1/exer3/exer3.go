// exer3 测量string + 和 join执行效率上的差异
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	s1, sep1 := "1", ""
	start1 := time.Now()

	for i := 1; i < 1000000; i++ {
		sep1 += s1
	}
	fmt.Println("add: ", time.Since(start1).Seconds())

	sep2 := ""
	s2 := []string{"1"}
	start2 := time.Now()

	for i := 1; i < 1000000; i++ {
		sep2 += strings.Join(s2, "")
	}
	fmt.Println("join: ", time.Since(start2).Seconds())
}
