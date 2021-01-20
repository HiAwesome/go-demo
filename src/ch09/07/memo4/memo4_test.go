//
// @author moqi
// On 2021/1/20 23:33:56
package memo4_test

import (
	"moqi.com/go/ch09/07/memo4"
	"moqi.com/go/ch09/07/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo4.New(httpGetBody)
	memotest.Sequential(m)
}

func TestConcurrent(t *testing.T) {
	m := memo4.New(httpGetBody)
	memotest.Concurrent(m)
}
