package main

import (
	"crypto/sha256"
	"fmt"
)

/**
 * 数组
 *
 * @author moqi
 * On 2021/1/2 22:40
 */
func main() {

	f1()

	f2()

	f3()

	f4()

	f5()

	f6()

}

func f6() {
	a := [32]byte{29: 1, 2, 3}
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1 2 3]
	fmt.Println(a)
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(zero(&a))
	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	fmt.Println(zeroFast(&a))
	fmt.Println()
}

func zero(ptr *[32]byte) [32]byte {
	for i := range ptr {
		ptr[i] = 0
	}
	return *ptr
}

func zeroFast(ptr *[32]byte) [32]byte {
	*ptr = [32]byte{}
	return *ptr
}

func f5() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println()
}

func f4() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	d := [3]int{1, 2}
	// Invalid operation: a == d (mismatched types [2]int and [3]int)
	// fmt.Println(a == d)
	fmt.Println(d)
	fmt.Println()
}

func f3() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(symbol)
	fmt.Println()

	r := [...]int{10: -1}
	fmt.Println(r)
	fmt.Println()
}

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func f2() {
	var q = [3]int{1, 2, 3}
	var r = [3]int{1, 2}
	fmt.Println(r[2], q[1])
	fmt.Println()

	var q1 = [...]int{1, 2, 3, 4}
	fmt.Println(q1)
	fmt.Printf("%T\n", q1)
	fmt.Println()
}

func f1() {
	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
	fmt.Println()
}
