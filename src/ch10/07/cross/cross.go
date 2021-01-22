//
// @author moqi
// On 2021/1/22 23:31:37
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
