// test of echo command
package main

import (
	"bytes"
	"fmt"
	"testing"
)

// 要注意的是测试代码和产品代码在同一个包。
// 虽然是 main 包，也有对应的 main 入口函数
// 但是在测试的时候 main 包只是 TestEcho 测试函数导入的一个普通包
// 里面 main 函数并没有被导出，而是被忽略的
func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
		{true, ",", []string{"a", "b", "c"}, "a b c\n"}, // NOTE: wrong expectation!
	}

	for _, t1 := range tests {
		desc := fmt.Sprintf("echo(%v, %q, %q)",
			t1.newline, t1.sep, t1.args)

		out = new(bytes.Buffer) // captured output
		if err := echo(t1.newline, t1.sep, t1.args); err != nil {
			t.Errorf("%s failed: %v", desc, err)
			continue
		}

		got := out.(*bytes.Buffer).String()
		if got != t1.want {
			t.Errorf("%s == %q, want %q", desc, got, t1.want)
		}
	}
}
