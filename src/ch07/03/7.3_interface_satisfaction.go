package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

/**
 * 实现接口的条件
 *
 * 有类的实现操作和实例的实现操作
 *
 * @author moqi
 * On 2021/1/10 21:55:36
 */
func main() {

	f1()

	f2()

}

func f2() {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = new(bytes.Buffer)
	any = map[string]int{"one": 1}
	fmt.Println(any)
	fmt.Println()
}

func f1() {
	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	// Cannot use 'time.Second' (type Duration) as type io.Writer Type does not implement
	// 'io.Writer' as some methods are missing: Write(p []byte) (n int, err error)
	// w = time.Second
	fmt.Println(w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout
	// Cannot use 'new(bytes.Buffer)' (type *bytes.Buffer) as type io.ReadWriteCloser Type does not implement
	// 'io.ReadWriteCloser' as some methods are missing: Close() error
	// rwc = new(bytes.Buffer)
	fmt.Println(rwc)
	fmt.Println()

	w = rwc
	// Cannot use 'w' (type io.Writer) as type io.ReadWriteCloser Type does not implement
	// 'io.ReadWriteCloser' as some methods are missing: Read(p []byte) (n int, err error) Close() error
	// rwc = w
}
