package main

import (
	"flag"
	"fmt"
	"strings"
)

/*
 * echo 更新
 *
 * go run a03_echo4.go a bc def
 * > a bc def
 *
 * go run a03_echo4.go -s / a bc def
 * > a/bc/def
 *
 * go run a03_echo4.go -n a bc def
 * > a bc def%
 *
 * go run a03_echo4.go -help
 * Usage of /var/folders/bq/d525bbp57776p5l17j_q12s80000gp/T/go-build437869626/b001/exe/a03_echo4:
 * -n	omit trailing newline
 * -s string
 *   	separator (default " ")
 *
 * @author moqi
 * On 12/14/20 11:26
 */
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
