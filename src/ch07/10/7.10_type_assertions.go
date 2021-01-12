package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/**
 * 类型断言
 *
 * @author moqi
 * On 2021/1/12 21:00:20
 */
func main() {
	var w io.Writer = os.Stdout
	f, ok := w.(*os.File)
	fmt.Println(f, ok)
	b, ok := w.(*bytes.Buffer)
	fmt.Println(b, ok)
}
