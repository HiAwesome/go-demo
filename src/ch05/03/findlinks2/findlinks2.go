package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

/**
 * 多返回值
 * find links 2
 *
 * @author moqi
 * On 2021/1/4 22:59
 */
// ~/Code/go-demo/src(main ✗) go run ch05/03/findlinks2/findlinks2.go http://golang.org
// https://support.eji.org/give/153413/#!/donation/checkout
// /
// /doc/
// /pkg/
// /project/
// /help/
// /blog/
// https://play.golang.org/
// /dl/
// https://tour.golang.org/
// https://blog.golang.org/
// /doc/copyright.html
// /doc/tos.html
// http://www.google.com/intl/en/policies/privacy/
// http://golang.org/issues/new?title=x/website:
// https://google.com
func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// findLinks performs an HTTP Get request for url
// parses the response as HTML, and extracts and returns the links.
// 按照惯例，函数的最后一个 bool 类型的返回值表示函数是否运行成功
func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
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

// 如果一个函数所有的返回值都有显式的变量名，那么该函数的 return 语句可以省略操作数，这称之为 bare return
// 按照返回值列表的次序，这个例子中每个 return 语句相当于 return words, images, err
// 当一个函数有多处 return 语句以及许多返回值时，bare return 可以减少代码的重复，但是使得代码难以被理解
// 不宜过度使用 bare return
// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	/* ... */
	return 0, 0
}
