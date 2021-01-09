package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

/**
 * 使用 defer 优化 resp.Body.Close() 的调用
 *
 * @author moqi
 * On 2021/1/9 22:59:53
 */
func main() {
	for _, arg := range os.Args[1:] {
		if err := title(arg); err != nil {
			fmt.Fprintf(os.Stderr, "title: %v\n", err)
		}
	}
}

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// 一条 defer 语句替代了之前的所有 resp.Body.Close()
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	const textHTML = "text/html"
	if ct != textHTML && !strings.HasPrefix(ct, textHTML) {
		return fmt.Errorf("%s has type %s, not %s", url, ct, textHTML)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	forEachNode(doc, visitNode, nil)
	return nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
