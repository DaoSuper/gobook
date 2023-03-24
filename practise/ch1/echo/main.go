package main

import (
	"fmt"
	"os"
)

func main() {
	pr1()
	pr2()
}

// 练习 1.1： 修改echo程序，使其能够打印os.Args[0]，即被执行命令本身的名字。
func pr1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("打印结果：", s)
}

// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
func pr2() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, ":", os.Args[i])
	}
}

// go run main.go 56 12 14
