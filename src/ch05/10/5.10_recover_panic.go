package main

/**
 * Recover 捕获异常
 *
 * @author moqi
 * On 2021/1/10 13:52:44
 */
func main() {
	f1()
}

// 练习 5.19：使用 panic 和 recover 编写一个不包含 return 语句但能返回一个非零值的函数
func f1() {
	defer func() {
		switch p := recover(); p {
		default:
			panic(100)
		}
	}()
}
