package main

import "fmt"

/**
 * 切片
 *
 * @author moqi
 * On 2021/1/2 23:03
 */
func main() {

	f1()

	f2()

	f3()

	f4()

	f5()

	f7()

	f8()

	f9()

}

func f9() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("s = ", s)
	s1 := remove(s, 1)
	fmt.Println("s1 = ", s1)
	fmt.Println()

	n := []int{5, 6, 7, 8, 9}
	fmt.Println("remove(n, 2) = ", remove(n, 2))
	fmt.Println()
}

// 移除某个位置元素，但是不需要保持顺序，用最后一个元素覆盖即可
func removeNotNeedOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// 移除 slice 中间某个位置元素，并将后面元素向前依次移动一位
func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func f8() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n", data)
	fmt.Println()
}

// 使用 append 实现 non-emtpy
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// Nonempty is an example of an in-place slice algorithm.
// nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func f7() {
	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	// append the slice x, 后面的 x 需要加三个点表示变长参数
	x = append(x, x...)
	fmt.Println("x = ", x)
	fmt.Println()
}

// slice like this
type IntSlice struct {
	ptr      *int
	len, cap int
}

func f5() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt1(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	fmt.Println()
}

// 三个点表示接收变长的参数为 slice
func appendInt1(x []int, y ...int) []int {
	var z []int
	zLen := len(x) + 1
	if zLen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zLen]
	} else {
		// There is insufficient space. Allocate a new Array.
		// Grow by doubling, for amortized linear complexity.
		zCap := zLen
		if zCap < 2*len(x) {
			zCap = 2 * len(x)
		}
		z = make([]int, zLen, zCap)
		copy(z, x) // a built-in function; see text
	}
	// ... expand z to at least zLen...
	copy(z[len(x):], y)
	// z[len(x)] = y
	return z
}

func f4() {
	var runes []rune
	for _, r := range "Hello, world" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	fmt.Println()
}

func f3() {
	var s []int
	fmt.Println("len(s) = ", len(s), " s == nil = ", s == nil)
	s = nil
	fmt.Println("len(s) = ", len(s), " s == nil = ", s == nil)
	s = []int(nil)
	fmt.Printf("%T\n", s)
	fmt.Println("len(s) = ", len(s), " s == nil = ", s == nil)
	s = []int{}
	fmt.Println("len(s) = ", len(s), " s == nil = ", s == nil)
	fmt.Println()
}

func equals(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func f2() {
	a := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("a = ", a)
	reverse(a[:])
	fmt.Println("a = ", a)
	fmt.Println()

	s := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("s = ", s)
	// Rotate s left by two positions
	reverse(s[:2])
	fmt.Println("s = ", s)
	reverse(s[2:])
	fmt.Println("s = ", s)
	reverse(s)
	fmt.Println("s = ", s)
	fmt.Println()
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func f1() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April",
		5: "May", 6: "June", 7: "July", 8: "August",
		9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)
	fmt.Println(months[0])
	fmt.Println()

	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)
	fmt.Println(summer)
	fmt.Println()

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appear in both\n", s)
			}
		}
	}
	fmt.Println()

	// fmt.Println(summer[:20])
	endlessSummer := summer[:5]
	fmt.Println("endlessSummer = ", endlessSummer)
	fmt.Println()
}
