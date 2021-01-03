package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

/**
 * 字符串
 *
 * @author moqi
 * On 2020/12/31 21:58
 */
func main() {
	f1()

	f2()

	f3()

	f4()

	f5()

	f6()

	f7()

	f8()

	f9()
}

func f9() {
	x, err1 := strconv.Atoi("123")
	fmt.Println(x, err1)

	y, err2 := strconv.ParseInt("123", 10, 64)
	fmt.Println(y, err2)
	fmt.Println()
}

func f8() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))
	fmt.Println()

	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println()

	s := fmt.Sprintf("x=%b", x)
	fmt.Println(s)
	fmt.Println()
}

func f7() {
	fmt.Println(intsToString([]int{1, 2, 3}))
	fmt.Println(fmt.Sprint([]int{1, 2, 3}))
	fmt.Println()
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		_, _ = fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func f6() {
	s := "abc"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s2)
	fmt.Println()
}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func f5() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
	fmt.Println()
}

// 第二个版本，使用 strings.LastIndex 库函数
// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func f4() {
	// "program" in Japanese katakana
	s := "プログラム"
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
	fmt.Println()

	fmt.Println(string(rune(65)))
	fmt.Println(string(rune(0x4eac)))
	// 如果对应码点的字符是无效的，则用 \uFFFD 无效字符作为替换
	fmt.Println(string(rune(1234567)))
	fmt.Println()
}

func f3() {
	s := "hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println()

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	fmt.Println()

	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}
}

func f2() {
	s := "left foot"
	t := s
	s += ", right foot"
	fmt.Println(s)
	fmt.Println(t)
	fmt.Println()
}

// 是否是前缀
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// 是否是后缀
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// 是否包含子串测试
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}

func f1() {
	s := "hello, world"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])
	fmt.Println()

	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[7:])
	fmt.Println(s[:])
	fmt.Println()

	fmt.Println("goodbye" + s[5:])
	fmt.Println()
}
