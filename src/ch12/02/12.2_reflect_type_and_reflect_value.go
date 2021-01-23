// reflect.Type 和 reflect.Value
// @author moqi
// On 2021/1/23 22:39:26
package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {

	f1()

	f2()

	f3()

}

func f3() {
	v := reflect.ValueOf(3)
	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)
	fmt.Println()
}

func f2() {
	v := reflect.ValueOf(3)
	fmt.Println(v)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	fmt.Println()

	t := v.Type()
	fmt.Println(t.String())
	fmt.Println()
}

func f1() {
	t := reflect.TypeOf(3)
	fmt.Println(t.String())
	fmt.Println(t)
	fmt.Println()

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	fmt.Println()

	fmt.Printf("%T\n", 3)
	fmt.Println()
}
