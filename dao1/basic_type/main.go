package main

import (
	"fmt"
	"math"
)

// 全局变量m
var m = 100

func main() {

	// var 变量名 类型 = 表达式

	//短变量声明
	n := 10
	m := 200
	fmt.Println("短变量声明", m, n)

	//匿名变量
	x, _ := foo()
	fmt.Println("匿名变量", x)

	fmt.Println("str := \"c:\\pprof\\main.exe\"")

	s1 := `第一行
    第二行
    第三行
    `
	fmt.Println(s1)

	traversalString()

	changeString()

	sqrtDemo()
}

func foo() (int, string) {
	return 10, "Q1mi"
}

// 遍历字符串
func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {
	s1 := "hello"
	// 强制类型转换
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
