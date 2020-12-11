package main

import (
	"fmt"
	"os"
)

/*
 * 打印出命令行参数 V4
 *
 * @author moqi
 * On 12/11/20 15:49
 */

func main() {
	fmt.Println(os.Args[1:])
}
