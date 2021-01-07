package word2

import "unicode"

// 判断字符串是否回文
// 采用 rune 序列，忽略空格和字母大小写
func IsPalindrome(s string) bool {
	// 基准测试优化2: 开始时为每个字符串预先分配一个足够大的数组，性能提升 50%
	letters := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	// 基准测试优化1: 只比较一次的优化，避免每个比较出现两次，性能提升 1%
	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}

	return true
}
