package main

import (
	"fmt"
	"time"
)

/**
 * Goroutines 示例: spinner
 *
 * 由于函数使用低效的递归，所以会运行相当长的时间，在此期间我们让用户看到一个可见的标识
 * 来表明程序依然在正常运行。
 *
 * @author moqi
 * On 2021/1/12 22:24:40
 */
func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFIbonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
