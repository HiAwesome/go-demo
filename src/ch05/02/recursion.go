package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

/**
 * 递归
 *
 * 使用 go build 完成之后将可执行文件放在 shell 下，然后在 shell 下运行
 * ~/Code/go-demo/shell(main ✗) ./fetch https://golang.org | ./recursion
 * https://support.eji.org/give/153413/#!/donation/checkout
 * /
 * /doc/
 * /pkg/
 * /project/
 * /help/
 * /blog/
 * https://play.golang.org/
 * /dl/
 * https://tour.golang.org/
 * https://blog.golang.org/
 * /doc/copyright.html
 * /doc/tos.html
 * http://www.google.com/intl/en/policies/privacy/
 * http://golang.org/issues/new?title=x/website:
 * https://google.com
 *
 * @author moqi
 * On 2021/1/4 21:31
 */
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
