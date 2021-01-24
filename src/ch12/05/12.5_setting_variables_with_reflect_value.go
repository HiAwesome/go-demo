// 通过 reflect.Value 修改值
// @author moqi
// On 2021/1/24 20:41:51
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	f1()

	f2()

	f3()

	f4()
}

func f4() {
	stdout := reflect.ValueOf(os.Stdout).Elem()
	fmt.Println(stdout.Type())
	fd := stdout.FieldByName("fd")
	// fmt.Println(fd.Int())
	// fd.SetInt(2)
	fmt.Println(fd.CanAddr(), fd.CanSet())
	fmt.Println()
}

func f3() {
	x := 1
	rx := reflect.ValueOf(&x).Elem()
	rx.SetInt(2)
	fmt.Println(x)
	rx.Set(reflect.ValueOf(3))
	fmt.Println(x)
	// rx.SetString("hello")
	// rx.Set(reflect.ValueOf("hello"))

	var y interface{}
	ry := reflect.ValueOf(&y).Elem()
	// ry.SetInt(2)
	ry.Set(reflect.ValueOf(3))
	fmt.Println(y)
	// ry.SetString("hello")
	ry.Set(reflect.ValueOf("hello"))
	fmt.Println(y)
	fmt.Println()
}

func f2() {
	x := 2
	d := reflect.ValueOf(&x).Elem()
	px := d.Addr().Interface().(*int)
	*px = 3
	fmt.Println(x)

	d.Set(reflect.ValueOf(4))
	fmt.Println(x)
	fmt.Println()
}

func f1() {
	x := 2
	a := reflect.ValueOf(2)
	b := reflect.ValueOf(x)
	c := reflect.ValueOf(&x)
	d := c.Elem()
	fmt.Println(a, b, c, d)
	fmt.Println(a.CanAddr())
	fmt.Println(b.CanAddr())
	fmt.Println(c.CanAddr())
	fmt.Println(d.CanAddr())
	fmt.Println()
}
