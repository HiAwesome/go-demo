package main

import "fmt"

/**
 * squares
 * 这个例子证明，函数值不仅仅是一串代码，还记录了状态。
 * 在 squares 中定义的匿名内部函数可以访问和更新 squares 中的局部变量，这意味着匿名函数和 squares 中，存在变量引用。
 * 这就是函数值属于引用类型和函数值不可比较的原因。
 * Go 使用闭包 (closures) 技术实现函数值，Go 程序员也把函数值叫做闭包。
 *
 * 通过这个例子，我们看到变量的生命周期不由它的作用域决定：squares 返回后，变量 x 仍然隐式的存在于 f 中。
 *
 * @author moqi
 * On 2021/1/9 20:13:34
 */
func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

// squares 返回一个匿名函数
// 该函数每次被调用时都会返回下一个数的平方
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
