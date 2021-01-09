package main

import (
	"fmt"
	"os"
)

/**
 * 可变参数
 *
 * @author moqi
 * On 2021/1/9 21:17:41
 */
func main() {

	f1()

	f2()

	f3()

}

func f3() {
	errorf(10, "%s %s", "wrong1", "wrong2")
	fmt.Println()
}

// 可变参数函数经常被用于格式化字符串
// 这个函数构造了一个以行号开头，经过格式化的错误信息
// 函数名的后缀 f 是一种通用的命名规范，代表该可变参数函数可以接收 Printf 风格的格式化字符串
// interface{} 表示函数的最后一个参数可以接收任意类型
func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func f2() {
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
	fmt.Println()
}

func f(v ...int) {}
func g(v []int)  {}

func f1() {
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	values := []int{1, 2, 3, 4}
	// Python 的解包
	fmt.Println(sum(values...))
	fmt.Println()
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
