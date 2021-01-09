package main

import (
	"fmt"
	"log"
	"moqi.com/go/ch05/06/links"
	"os"
)

/**
 * 模拟一个网络爬虫，使用广度优先算法
 *
 * @author moqi
 * On 2021/1/9 20:51:49
 */
func main() {
	breadthFirst(crawl, os.Args[1:])
}

// breadthFirst calls f for each item in the worklist.
//  ANy items returned by f are added to the worklist.
// f is called at most once for each item
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				// append 的参数 f(item)... 会将 f 返回的一组元素一个个添加到 worklist 中
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
