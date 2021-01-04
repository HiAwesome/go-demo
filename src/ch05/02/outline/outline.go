package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

/**
 * outline
 *
 * @author moqi
 * On 2021/1/4 22:27
 */
// ~/Code/go-demo/shell(main ✗) ./fetch https://golang.org | ./outline
// ...
// [html body]
// [html body header]
// [html body header div]
// [html body footer div ul li]
// [html body footer div ul li a]
// [html body footer div a]
// [html body script]
// ...
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

// 我们通过递归的方式遍历整个 HTML 节点树，并输出树的结构。
// 在 outline 内部，每遇到一个 HTML 元素标签，就将其入栈，并输出。
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
