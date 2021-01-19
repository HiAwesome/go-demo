// 内存同步: 运行多次，观察差异
// @author moqi
// On 2021/1/19 22:19:04
package main

import "fmt"

func main() {

	defer fmt.Println("Everything OK")

	var x, y int
	go func() {
		x = 1                   // A1
		fmt.Print("y:", y, " ") // A2
	}()
	go func() {
		y = 1                   // B1
		fmt.Print("x:", x, " ") // B2
	}()

}
