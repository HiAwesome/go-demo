// Countdown implements the countdown for a rocket launch.
// @author moqi
// On 2021/1/17 23:08:15
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
