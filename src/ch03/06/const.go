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

	f3()

	f4()

	f5()

	f6()

}

func f6() {
	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')
	fmt.Println()
}

func f5() {
	var f float64 = 212
	// "100", (f - 32) * 5 is a float64
	fmt.Println((f - 32) * 5 / 9)
	// "0"; 5/9 is an untyped integer, 0
	fmt.Println(5 / 9 * (f - 32))
	// "100"; 5.0/9.0 is an untyped float
	fmt.Println(5.0 / 9.0 * (f - 32))
	fmt.Println()
}

func f4() {
	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
		ZiB
		YiB
	)

	fmt.Println(KiB, MiB, GiB, TiB)
	fmt.Println(PiB, EiB, float64(ZiB), float64(YiB))
	fmt.Println()
}

func f3() {
	var v = net.FlagMulticast | net.FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v))
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))
	fmt.Printf("%b %t\n", v, IsCast(v))
	fmt.Println()
}

func IsUp(v net.Flags) bool {
	return v&net.FlagUp == net.FlagUp
}

func TurnDown(v *net.Flags) {
	*v &^= net.FlagUp
}

func SetBroadcast(v *net.Flags) {
	*v |= net.FlagBroadcast
}

func IsCast(v net.Flags) bool {
	return v&(net.FlagBroadcast|net.FlagMulticast) != 0
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
