package main

import (
	"fmt"
	"math"
)

/**
 * 函数声明
 *
 * @author moqi
 * On 2021/1/4 21:16
 */
func main() {

	fmt.Println(hypot(3, 4))
	fmt.Println()

	// 四种方法声明拥有两个 int 型参数和一个 int 型返回值的函数
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
	fmt.Println()

	// 没有函数体的函数声明，表示该函数不是以 Go 实现的
	// 例如在 sin.go 中有 func Sin(x float64) float64

}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func f1(i, j, k int, s, t string) {
	/* ... */
}

func f2(i int, j int, k int, s string, t string) {
	/* ... */
}
