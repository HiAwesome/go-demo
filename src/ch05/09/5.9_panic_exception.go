package main

import (
	"fmt"
	"os"
	"runtime"
)

/**
 * Panic 异常
 *
 * @author moqi
 * On 2021/1/10 00:03:53
 */
func main() {

	// 将 panic 机制类比其他语言异常机制的读者可能会惊讶，runtime.Stack 为何能输出已经被释放函数的信息？
	// 因为在 Go 的 panic 机制中，延迟函数的调用在释放堆栈信息之前。
	defer printStack()
	f(3)

}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
