package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

/*
 * 查找重复的行 V3
 *
 * @author moqi
 * On 12/11/20 16:49
 */

func main() {

	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
