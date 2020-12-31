package main

import (
	"fmt"
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

	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
	fmt.Println()

}

// 第一个版本，全部手工硬编码实现
// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
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
