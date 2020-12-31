package main

/**
 * @author moqi
 * On 2020/12/31 21:41
 */
func main() {
	print(btoi(true))
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}
