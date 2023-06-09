// 一个处理输入的循环，在每个元素上执行计算处理，在处理的同时或最后产生输出
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dup1()
}

// 打印标准输入中多次出现的行，以重复次数开头
func dup1() {
	counts := make(map[string]int)
	//Scanner类型是bufio包最有用的特性之一，它读取输入并将其拆成行或单词
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// Printf
// %d          十进制整数
// %x, %o, %b  十六进制，八进制，二进制整数。
// %f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
// %t          布尔：true或false
// %c          字符（rune） (Unicode码点)
// %s          字符串
// %q          带双引号的字符串"abc"或带单引号的字符'c'
// %v          变量的自然形式（natural format）
// %T          变量的类型
// %%          字面上的百分号标志（无操作数）
