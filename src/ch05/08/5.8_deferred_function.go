package main

import "fmt"

/**
 * Deferred 函数
 *
 * 你只需要在调用普通函数或方法前加上关键字 defer，就完成了 defer 所需要的语法。
 * 当执行到该语句时，函数和参数表达式得到计算，但直到包含该 defer 语句的函数执行完毕时，defer 后的函数才会被执行。
 * 不论包含 defer 语句的函数是通过 return 正常结束，还是由于 panic 导致的异常结束。
 * 你可以在一个函数中执行多条 defer 语句，他们的执行顺序与声明顺序相反。
 *
 * defer 语句经常被用于处理成对的操作，如打开/关闭，连接/断开连接，加锁/释放锁。
 * 通过 defer 机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。
 * 释放资源的 defer 应该直接跟在请求资源的语句后面。
 *
 * 调试复杂程序时，defer 机制也常被用于记录何时进入和退出函数。
 *
 * @author moqi
 * On 2021/1/9 22:36:44
 */
func main() {

	f1()

}

func f1() {
	_ = double(4)
	fmt.Println()

	fmt.Println(triple(4))
	fmt.Println()
}

// 对匿名函数采用 defer 机制，可以使其观察函数的返回值
func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	}()

	return x + x
}

// 被延迟执行的匿名函数甚至可以修改函数返回给调用者的返回值
func triple(x int) (result int) {
	defer func() {
		result += x
	}()
	return double(x)
}
