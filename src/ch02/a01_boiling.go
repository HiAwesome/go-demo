package main

import "fmt"

/*
 * 声明
 *
 * @author moqi
 * On 12/14/20 10:48
 */
const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	// Output: boiling point = 212°F or 100°C
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}
