//
// @author moqi
// On 2021/1/20 22:58:31
package memo2_test

import (
	"moqi.com/go/ch09/07/memo2"
	"moqi.com/go/ch09/07/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo2.New(httpGetBody)
	memotest.Sequential(m)
}

func TestConcurrent(t *testing.T) {
	m := memo2.New(httpGetBody)
	memotest.Concurrent(m)
}
