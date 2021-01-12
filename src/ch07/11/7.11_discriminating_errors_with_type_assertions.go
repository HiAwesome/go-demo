package main

import (
	"fmt"
	"os"
)

/**
 * 基于类型断言区别错误类型
 *
 * @author moqi
 * On 2021/1/12 21:34:40
 */
func main() {

	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Println("%#v\n", err)
	fmt.Println(os.IsNotExist(err))

}
