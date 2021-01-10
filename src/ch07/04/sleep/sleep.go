// The sleep program sleeps for a specified period of time.
package main

import (
	"flag"
	"fmt"
	"time"
)

/**
 *
 *
 * @author moqi
 * On 2021/1/10 22:09:55
 */
func main() {
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

var period = flag.Duration("period", 1*time.Second, "sleep period")
