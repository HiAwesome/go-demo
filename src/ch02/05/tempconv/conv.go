package tempconv

/*
 * 包和文件
 * 转换函数
 *
 * @author moqi
 * On 12/14/20 14:51
 */
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
