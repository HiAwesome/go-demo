//
// @author moqi
// On 2021/1/19 22:41:06
package memo1_test

import (
	"moqi.com/go/ch09/07/memo1"
	"moqi.com/go/ch09/07/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo1.New(httpGetBody)
	memotest.Sequential(m)
}

func TestConcurrent(t *testing.T) {
	m := memo1.New(httpGetBody)
	memotest.Concurrent(m)
}

// ~/Code/go-demo/src/ch09/07/memo1(main ✗) go test -run=TestConcurrent -race -v
// === RUN   TestConcurrent
// 2021/01/19 22:51:07 Get "https://play.golang.org": EOF
// https://godoc.org, 1.355283643s, 7480 bytes
// https://godoc.org, 2.109531661s, 7480 bytes
// 2021/01/19 22:51:13 Get "https://golang.org": EOF
// 2021/01/19 22:51:16 Get "https://golang.org": net/http: TLS handshake timeout
// ==================
// WARNING: DATA RACE
// Write at 0x00c000114d80 by goroutine 29:
//  runtime.mapassign_faststr()
//      /Users/moqi/.gvm/gos/go1.15.6/src/runtime/map_faststr.go:202 +0x0
//  moqi.com/go/ch09/07/memo1.(*Memo).Get()
//      /Users/moqi/Code/go-demo/src/ch09/07/memo1/memo1.go:26 +0x1cd
//  moqi.com/go/ch09/07/memotest.Concurrent.func1()
//      /Users/moqi/Code/go-demo/src/ch09/07/memotest/memotest.go:72 +0xe5
//
// Previous write at 0x00c000114d80 by goroutine 9:
//  runtime.mapassign_faststr()
//      /Users/moqi/.gvm/gos/go1.15.6/src/runtime/map_faststr.go:202 +0x0
//  moqi.com/go/ch09/07/memo1.(*Memo).Get()
//      /Users/moqi/Code/go-demo/src/ch09/07/memo1/memo1.go:26 +0x1cd
//  moqi.com/go/ch09/07/memotest.Concurrent.func1()
//      /Users/moqi/Code/go-demo/src/ch09/07/memotest/memotest.go:72 +0xe5
//
// Goroutine 29 (running) created at:
//  moqi.com/go/ch09/07/memotest.Concurrent()
//      /Users/moqi/Code/go-demo/src/ch09/07/memotest/memotest.go:69 +0x10c
//  moqi.com/go/ch09/07/memo1_test.TestConcurrent()
//      /Users/moqi/Code/go-demo/src/ch09/07/memo1/memo1_test.go:21 +0xdd
//  testing.tRunner()
//      /Users/moqi/.gvm/gos/go1.15.6/src/testing/testing.go:1123 +0x202
//
// Goroutine 9 (finished) created at:
//  moqi.com/go/ch09/07/memotest.Concurrent()
//      /Users/moqi/Code/go-demo/src/ch09/07/memotest/memotest.go:69 +0x10c
//  moqi.com/go/ch09/07/memo1_test.TestConcurrent()
//      /Users/moqi/Code/go-demo/src/ch09/07/memo1/memo1_test.go:21 +0xdd
//  testing.tRunner()
//      /Users/moqi/.gvm/gos/go1.15.6/src/testing/testing.go:1123 +0x202
// ==================
// 2021/01/19 22:51:16 Get "https://play.golang.org": net/http: TLS handshake timeout
// 2021/01/19 22:51:22 Get "http://www.gopl.io/": EOF
// 2021/01/19 22:51:24 Get "http://www.gopl.io/": EOF
//    testing.go:1038: race detected during execution of test
// --- FAIL: TestConcurrent (18.30s)
// === CONT
//    testing.go:1038: race detected during execution of test
// FAIL
// exit status 1
// FAIL	moqi.com/go/ch09/07/memo1	18.874s
