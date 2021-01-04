package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * 查找重复的行 V1
 *
 * @author moqi
 * On 12/11/20 15:52
 */

func main() {

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	// Note: ignoring potential errors from input.Err()

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
