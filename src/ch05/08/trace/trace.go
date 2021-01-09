package main

import (
	"log"
	"time"
)

/**
 * 2021/01/09 23:35:17 enter bigSlowOperation
 * 2021/01/09 23:35:27 exit bigSlowOperation (10.005183595s)
 *
 * @author moqi
 * On 2021/1/9 23:28:20
 */
func main() {
	bigSlowOperation()
}

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work...
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
