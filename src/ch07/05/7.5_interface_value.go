package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/**
 * 接口值
 *
 * 警告：一个不包含任何值的 nil 接口值和一个刚好包含 nil 指针的接口值是不同的。
 * 这个细微区别产生了一个容易绊倒每个 Go 程序员的陷阱。
 *
 * @author moqi
 * On 2021/1/10 22:34:49
 */
func main() {
	var w io.Writer
	// 在 fmt 内部，使用反射来获取接口动态类型的名称
	fmt.Printf("%T\n", w)
	w = os.Stdout
	fmt.Printf("%T\n", w)
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w)
	fmt.Println()
}
