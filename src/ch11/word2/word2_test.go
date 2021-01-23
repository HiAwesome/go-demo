package word2

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 表格化测试
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}

// 随机测试
// randomPalindrome returns a palindrome whose length and contents
// are derived from the predudo-random number generator rng.
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random length up to 24
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune up to '\u0999'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// Initialize a preudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("ISPalindrome(%q) = false", p)
		}
	}
}

// 基准测试
// go test -bench=.
// 和普通测试不同的是，默认情况下不运行任何基准测试。
// 我们通过 -bench 命令行标志参数手工指定要运行的基准测试函数
// 该参数是一个正则表达式，用于匹配要执行的基准测试函数的名字，默认值是空的
// 其中 . 模式可以匹配所有基准测试函数，因为这里只有一个基准测试函数
// 因此和 -bench=IsPalindrome 参数是等价的效果
//
// 快的程序往往是伴随较少的内存分配。 -benchmem 命令行标志参数将在报告中包含内存的分配数据统计
// 用一次内存分配代替多次内存分配节省了 75% 的分配调用次数和将近一半的内存需求
func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan, a canal: Panama")
	}
}

// 示例函数
// 三个用处：作为文档、测试时自身运行、提供一个真实的演练场
// 具体参考 https://blog.golang.org/examples
// 所以 Example 一般需要写上 Output 语句
// 典型的例外的情况是代码需要访问网络、或者 map 无法保证输出顺序时，不写 Output 语句
func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}
