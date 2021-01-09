package main

import (
	"fmt"
	"sort"
)

/**
 * 拓扑排序
 *
 * @author moqi
 * On 2021/1/9 20:35:31
 */
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// 这段代码用深度优先搜素了整张图，获得了符合要求的课程序列
func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	// 当匿名函数需要被递归调用时，我们必须首先声明一个变量，再将匿名函数赋值给这个变量
	// 如果不分成两步，函数字面量无法与 visitAll 绑定，我们也无法递归调用该匿名函数
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// prereqs 记录了每个课程的前置课程
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}
