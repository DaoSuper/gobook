// "_"标识符，用来忽略函数的某个返回值
package main

import (
	"fmt"
	"math"
)

func main() {
	x, _ := test()
	println(x)

	println("多返回值作为实参： ", add(test()))

	//匿名函数由一个不带函数名的函数声明和函数体组成
	getSqrt := func(a float64) float64 {
		return math.Sqrt(a)
	}

	fmt.Println("匿名函数", getSqrt(4))
}

func test() (int, int) {
	return 1, 2
}

func add(x, y int) int {
	return x + y
}
