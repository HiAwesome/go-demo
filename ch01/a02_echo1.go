package main

/*
 * 打印出命令行参数
 *
 * @author moqi
 * On 12/11/20 15:12
 */

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
