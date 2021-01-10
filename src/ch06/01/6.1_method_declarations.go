package main

import (
	"fmt"
	"moqi.com/go/ch06/01/geometry"
)

/**
 * 方法声明
 *
 * @author moqi
 * On 2021/1/10 14:17:57
 */
func main() {

	f1()

	f2()

}

func f2() {
	perim := geometry.Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	// 独立的函数在外部调用需要写包名才可以，而类型的方法却不需要
	fmt.Println(geometry.PathDistance(perim)) // standalone function
	fmt.Println(perim.Distance())             // method of geometry.Path
	fmt.Println()
}

func f1() {
	p := geometry.Point{1, 2}
	q := geometry.Point{4, 6}
	fmt.Println(geometry.Distance(p, q)) // function call
	fmt.Println(p.Distance(q))           // method call
	fmt.Println()
}
