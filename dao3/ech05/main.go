// 递归，就是在运行的过程中调用自
// 1.子问题须与原始问题为同样的事，且更为简单。
// 2.不能无限制地调用本身，须有个出口，化简为非递归状况处理。
package main

import "fmt"

func main() {
	var i int = 7
	fmt.Printf("Factorial of %d is %d\n", i, factorial(i))

	for i = 0; i < 10; i++ {
		fmt.Printf("fibonaci %d\n", fibonaci(i))
	}
}

// 数字阶乘:一个正整数的阶乘（factorial）是所有小于及等于该数的正整数的积，并且0的阶乘为1。自然数n的阶乘写作n!
func factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * factorial((i - 1))
}

// 斐波那契数列(Fibonacci)
// 这个数列从第3项开始，每一项都等于前两项之和。
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}

	if i == 1 {
		return 1
	}

	return factorial(i-1) + factorial(i-2)
}
