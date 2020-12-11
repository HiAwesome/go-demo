package main

/*
 * 打印出命令行参数 V2
 *
 * @author moqi
 * On 12/11/20 15:34
 */

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)
}
