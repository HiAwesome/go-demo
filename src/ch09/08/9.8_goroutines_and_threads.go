// Goroutines 和线程
// @author moqi
// On 2021/1/20 23:58:44
package main

import "fmt"

// $ GOMAXPROCS=1 go run hacker-cliché.go
// 111111111111111111110000000000000000000011111...
//
// $ GOMAXPROCS=2 go run hacker-cliché.go
// 010101010101010101011001100101011010010100110...
func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
