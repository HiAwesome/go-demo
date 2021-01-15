// 不带缓存的 Channels
// @author moqi
// On 2021/1/15 22:29:14
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close()
	// go mustCopy(os.Stdout, conn)
	// mustCopy(conn, os.Stdin)
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		// 有些消息事件并不携带额外的信息，它仅仅是用作两个 goroutine 之间的同步
		// 这时候我们可以用 struct{} 空结构体作为 channels 元素的类型
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done // wait for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
