package main

/**
 * 布尔值
 *
 * @author moqi
 * On 2020/12/31 21:41
 */
func main() {
	println(btoi(true))
	println(itob(100))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}
