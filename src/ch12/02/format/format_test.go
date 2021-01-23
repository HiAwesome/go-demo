//
// @author moqi
// On 2021/1/23 23:10:51
package format_test

import (
	"fmt"
	"moqi.com/go/ch12/02/format"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var x int64 = 1
	var d = 1 * time.Nanosecond
	fmt.Println(format.Any(x))
	fmt.Println(format.Any(d))
	fmt.Println(format.Any([]int64{x}))
	fmt.Println(format.Any([]time.Duration{d}))
}
