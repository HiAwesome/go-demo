//
// @author moqi
// On 2021/1/23 23:10:51
package format

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	var x int64 = 1
	var d = 1 * time.Nanosecond
	fmt.Println(Any(x))
	fmt.Println(Any(d))
	fmt.Println(Any([]int64{x}))
	fmt.Println(Any([]time.Duration{d}))
}
