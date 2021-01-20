//
// @author moqi
// On 2021/1/20 23:17:40
package memo3_test

import (
	"moqi.com/go/ch09/07/memo3"
	"moqi.com/go/ch09/07/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo3.New(httpGetBody)
	memotest.Sequential(m)
}

func TestConcurrent(t *testing.T) {
	m := memo3.New(httpGetBody)
	memotest.Concurrent(m)
}
