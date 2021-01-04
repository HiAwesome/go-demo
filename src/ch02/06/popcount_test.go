package popcount_test

import (
	"moqi.com/go/ch02/06"
	"testing"
)

/*
 * 对 PopCount 的基准测试
 *
 * ~/Code/go-demo/src/ch02/06(main ✗) go test -bench=.
 * goos: darwin
 * goarch: amd64
 * pkg: moqi.com/go/ch02/06
 * BenchmarkPopCount-8                     1000000000               0.269 ns/op
 * BenchmarkBitCount-8                     1000000000               0.277 ns/op
 * BenchmarkPopCountByClearing-8           83732410                14.4 ns/op
 * BenchmarkPopCountByShifting-8           26011123                44.3 ns/op
 * PASS
 * ok      moqi.com/go/ch02/06     3.633s
 *
 * @author moqi
 * On 12/14/20 16:08
 */
func BitCount(x uint64) int {
	// Hacker's Delight, Figure 5-2.
	x = x - ((x >> 1) & 0x5555555555555555)
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)
	return int(x & 0x7f)
}

func PopCountByClearing(x uint64) int {
	n := 0
	for x != 0 {
		x = x & (x - 1) // clear rightmost non-zero bit
		n++
	}
	return n
}

func PopCountByShifting(x uint64) int {
	n := 0
	for i := uint(0); i < 64; i++ {
		if x&(1<<i) != 0 {
			n++
		}
	}
	return n
}

// -- Benchmarks --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BitCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(0x1234567890ABCDEF)
	}
}
