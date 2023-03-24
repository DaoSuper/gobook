// 定义了程序的各种实体对象以及部分或全部的属性
// var(变量) 、const(常量)、type(类型)和func(函数实体对象)声明
package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)

	const freezingF = 32.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
