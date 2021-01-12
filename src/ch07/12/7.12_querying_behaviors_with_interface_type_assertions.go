package main

import "fmt"

/**
 * 通过类型断言询问行为
 *
 * @author moqi
 * On 2021/1/12 21:41:41
 */
func main() {
	a := formatOneValue("tom smith")
	fmt.Println(a)
}

func formatOneValue(x interface{}) string {
	if err, ok := x.(error); ok {
		return err.Error()
	}
	if str, ok := x.(fmt.Stringer); ok {
		return str.String()
	}
	// ... all other types ...
	return ""
}
