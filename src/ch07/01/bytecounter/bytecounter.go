// Byte Counter demonstrates an implementation of io.Writer that counts bytes.
package main

import "fmt"

/**
 * Byte Counter
 *
 * @author moqi
 * On 2021/1/10 21:48:47
 */
func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // 5

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // 12
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}
