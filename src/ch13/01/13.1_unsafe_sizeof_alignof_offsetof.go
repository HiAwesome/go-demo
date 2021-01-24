// unsafe.Sizeof, alignof 和 Offsetof
// @author moqi
// On 2021/1/24 23:08:14
package main

import (
	"fmt"
	"unsafe"
)

var x struct {
	a bool
	b int16
	c []int
}

func main() {

	fmt.Println(unsafe.Sizeof(float64(0)))
	fmt.Println()

	fmt.Println(unsafe.Sizeof(x), unsafe.Alignof(x))
	fmt.Println(unsafe.Sizeof(x.a), unsafe.Alignof(x.a), unsafe.Offsetof(x.a))
	fmt.Println(unsafe.Sizeof(x.b), unsafe.Alignof(x.b), unsafe.Offsetof(x.b))
	fmt.Println(unsafe.Sizeof(x.c), unsafe.Alignof(x.c), unsafe.Offsetof(x.c))
	fmt.Println()

}
