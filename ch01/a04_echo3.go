package main

import (
	"fmt"
	"os"
	"strings"
)

/*
 * 打印出命令行参数 V3
 *
 * @author moqi
 * On 12/11/20 15:45
 */

func main() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}
