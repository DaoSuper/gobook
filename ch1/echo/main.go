package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1()
	echo2()
	echo3()
}

// os包以跨平台的方式，提供了一些与操作系统交互的函数和变量
// os.Args变量是一个字符串（string）的切片（slice）
// os.Args的第一个元素：os.Args[0]，是命令本身的名字；其它的元素则是程序启动时传给它的参数
func echo1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("echo1: ", s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println("echo2: ", s)
}

// strings包的Join函数
func echo3() {
	fmt.Println("echo3: ", strings.Join(os.Args[1:], ""))
}

//go run main.go args测试
