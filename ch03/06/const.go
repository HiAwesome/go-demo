package main

import (
	"fmt"
	"net"
	"time"
)

/**
 * 常量
 *
 * @author moqi
 * On 2021/1/2 16:24
 */
const pi = 3.1415926 // approximately, math.Pi is a better approximation

const (
	e   = 2.71828182845904523536028747135266249775724709369995957496696763
	pi2 = 3.14159265358979323846264338327950288419716939937510582097494459
)

const IPv4Len = 4

// parseIpv4 parses an IPv4 address (d.d.d.d)
func parseIpv4(s string) net.IP {
	var p [IPv4Len]byte
	fmt.Println(len(p))
	return nil
}

func main() {

	f1()

	f2()

}

func f2() {
	const (
		a = 1
		b
		c = 2
		d
	)

	fmt.Println(a, b, c, d)
	fmt.Println()
}

func f1() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Println()
}
