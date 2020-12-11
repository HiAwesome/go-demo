package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

/*
 * 获取 URL
 *
 * @author moqi
 * On 12/11/20 17:26
 */

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := ioutil.ReadAll(resp.Body)
		_ = resp.Body.Close()

		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s", b)
	}
}
