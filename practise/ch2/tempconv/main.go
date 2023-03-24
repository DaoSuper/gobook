// 练习 2.1： 向tempconv包添加类型、常量和函数用来处理Kelvin绝对温度的转换，
// Kelvin 绝对零度是−273.15°C，Kelvin绝对温度1K和摄氏度1°C的单位间隔是一样的。
package main

import (
	"fmt"
	"strconv"
)

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15 // 绝对零度
	FreezingC     Celsius = 0       // 结冰点温度
	BoilingC      Celsius = 100     // 沸水温度
)

func (c Celsius) String() string { return fmt.Sprintf("%.3f°C", c) }

func (f Fahrenheit) String() string { return fmt.Sprintf("%.3f°F", f) }

// 摄氏度转化成华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// 华氏度转化成摄氏度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func Kelvin(f Fahrenheit) Celsius { return Celsius(f) + AbsoluteZeroC }

func main() {
	str, _ := strconv.ParseFloat("60", 10)
	fmt.Println("单位转换后的结果是:", Kelvin(Fahrenheit(str)))
}
