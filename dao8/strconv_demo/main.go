package main

import (
	"fmt"
	"strconv"
)

//strconv包实现了基本数据类型与其字符串表示的转换，主要有以下常用函数： Atoi()、Itia()、parse系列、format系列、append系列。
//https://golang.org/pkg/strconv/

func main() {
	//将字符串类型的整数转换为int类型
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}

	//Itoa()函数用于将int类型数据转换为对应的字符串表示
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2)

	b, _ := strconv.ParseBool("true")
	fmt.Printf("ParseBool type:%T value:%#v\n", b, b)

	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Printf("ParseFloat type:%T value:%#v\n", f, f)

	//返回字符串表示的整数值,接受正负号
	i, _ := strconv.ParseInt("-2", 10, 64)
	fmt.Printf("ParseInt type:%T value:%#v\n", i, i)

	//返回字符串表示的整数值
	u, _ := strconv.ParseUint("2", 10, 64)
	fmt.Printf("ParseUint type:%T value:%#v\n", u, u)

	// 返回字符串表示的bool值
	formatBool := strconv.FormatBool(true)
	fmt.Printf("formatBool type:%T value:%#v\n", formatBool, formatBool)

	formatFloat := strconv.FormatFloat(3.1415, 'E', -1, 64)
	fmt.Printf("formatFloat type:%T value:%#v\n", formatFloat, formatFloat)

	//返回字符串表示的整数值，接受正负号
	formatInt := strconv.FormatInt(-2, 16)
	fmt.Printf("formatInt type:%T value:%#v\n", formatInt, formatInt)

	//ParseUint类似ParseInt但不接受正负号，用于无符号整型。
	formatUint := strconv.FormatUint(2, 16)
	fmt.Printf("formatUint type:%T value:%#v\n", formatUint, formatUint)

}
