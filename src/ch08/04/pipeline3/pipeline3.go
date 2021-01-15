// 单方向的 Channel
// @author moqi
// On 2021/1/15 22:52:54
package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

// 仅发送类型 chan<- int
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

// 仅接收类型 <-chan int
func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
