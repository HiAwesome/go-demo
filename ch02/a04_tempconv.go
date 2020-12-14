package main

import "fmt"

/*
 * 类型声明示例
 *
 * @author moqi
 * On 12/14/20 14:34
 */
type Celsius float64    // 摄氏温度
type Fahrenheit float64 // 华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 声明一个 Celsius 类型的名称叫 String 的方法
// 该方法返回该类型对象 c 带着温度单位的字符串
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	// fmt.Println(c == f) compile error: type mismatch
	fmt.Println(c == Celsius(f))
	fmt.Println()

	c = FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Println(c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
	fmt.Println()
}
