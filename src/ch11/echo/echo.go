// echo prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout // modified during testing

var (
	n = flag.Bool("n", false, "omit trailing newline")
	s = flag.String("s", " ", "separator")
)

func main() {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, sep string, args []string) error {
	_, _ = fmt.Fprint(out, strings.Join(args, sep))
	if newline {
		_, _ = fmt.Fprintln(out)
	}
	return nil
}
