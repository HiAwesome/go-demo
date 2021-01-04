package main

import (
	"fmt"
	"moqi.com/go/ch02/05/tempconv"
	"os"
	"strconv"
)

/*
 * 包和文件
 *
 * @author moqi
 * On 12/14/20 15:48
 */

func main() {

	// f1()

	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}

}

func f1() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println()

	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	fmt.Println()
}
