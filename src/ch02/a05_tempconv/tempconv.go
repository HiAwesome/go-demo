package a05_tempconv

import "fmt"

/*
 * 包和文件
 * 变量、常量、方法
 *
 * @author moqi
 * On 12/14/20 14:49
 */
type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
