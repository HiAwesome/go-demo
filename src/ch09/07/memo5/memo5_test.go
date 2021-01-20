//
// @author moqi
// On 2021/1/20 23:52:34
package memo5_test

import (
	"moqi.com/go/ch09/07/memo5"
	"moqi.com/go/ch09/07/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo5.New(httpGetBody)
	memotest.Sequential(m)
}

func TestConcurrent(t *testing.T) {
	m := memo5.New(httpGetBody)
	memotest.Concurrent(m)
}
