// unsafe.Pointer
// @author moqi
// On 2021/1/24 23:15:37
package main

import (
	"fmt"
	"unsafe"
)

func main() {

	fmt.Printf("%#016x\n", Float64bits(1.0))

}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
