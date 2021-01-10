// Temp flag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"
	"moqi.com/go/ch07/04/tempconv"
)

/**
 * Temp flag
 *
 * @author moqi
 * On 2021/1/10 22:23:44
 */
func main() {
	flag.Parse()
	fmt.Println(*temp)
}

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")
