package main

import (
	"fmt"
	"strings"
)

/**
 * 函数值在 Go 中是一等公民
 *
 * @author moqi
 * On 2021/1/9 19:42:34
 */
func main() {

	f1()

	// f2()

	f3()

}

func f3() {
	fmt.Println(strings.Map(add1, "HAL-9000"))
	fmt.Println(strings.Map(add1, "VMS"))
	fmt.Println(strings.Map(add1, "Admix"))
	// 使用匿名函数
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "HAL-9000"))
}

func add1(r rune) rune {
	return r + 1
}

func f2() {
	var f func(int) int

	// Function call 'f(3)' may lead to nil pointer dereference
	f(3)

	if f != nil {
		f(3)
	}
}

func f1() {
	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)

	// Cannot use 'product' (type func(m int, n int) int) as type func(n int) int
	// f = product
}

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}
