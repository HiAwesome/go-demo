package main

import "os"

/**
 * 匿名函数
 * 警告：捕获迭代变量 的练习
 *
 * @author moqi
 * On 2021/1/9 20:12:45
 */

func main() {

	// rightWay()

	// badWay()

}

func badWay() {
	TempDirs := map[int]string{
		1: "someDir1",
		2: "someDir2",
		3: "someDir3",
	}

	var rmdirs []func()
	for _, dir := range TempDirs {
		os.MkdirAll(dir, 0755)
		// 因为函数值中记录的是循环变量的内存地址，而不是循环变量某一次的值。
		// 以 dir 为例，后续的迭代会不断更新 dir 的值，当删除操作执行时，for 循环已经完成
		// dir 中存储的值等于最后一次迭代的值。
		// 这意味着，每次对 os.RemoveAll 的调用删除的都是相同的目录
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir) // NOTE: incorrect!
		})
	}
	// ...do some work...
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func rightWay() {
	TempDirs := map[int]string{
		1: "someDir1",
		2: "someDir2",
		3: "someDir3",
	}

	var rmdirs []func()
	// 为了解决这个问题，我们引入一个与循环变量同名的局部变量，作为循环电量的副本。
	// 比如 dir := dir，虽然看起来很奇怪，但却很有用。
	for _, d := range TempDirs {
		// dir := dir // declares inner dir, initialized to out dir
		dir := d               // NOTE: necessary!
		os.MkdirAll(dir, 0755) // creates parent directories too
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}
	// ...do some work...
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}

func rightWayUseI() {
	TempDirs := map[int]string{
		1: "someDir1",
		2: "someDir2",
		3: "someDir3",
	}
	var rmdirs []func()
	dirs := TempDirs

	// 这个问题不仅存在于基于 range 的循环
	// 这个例子示范了对循环变量 i 的使用也存在同样的问题
	for i := 0; i < len(dirs); i++ {
		os.MkdirAll(dirs[i], 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dirs[i]) // NOTE: incorrect!
		})
	}

	// ...do some work...
	for _, rmdir := range rmdirs {
		rmdir() // clean up
	}
}
