//形参 函数定义时的参数
//实参 函数调用时所传递的变量

// 值传递(值的拷贝)：指在调用函数时将实际参数复制一份传递到函数中，不会影响到实际参数
// 引用传递(地址的拷贝)：是指在调用函数时将实际参数的地址传递到函数中， 会影响到实际参数
// map、slice、chan、指针、interface默认以引用的方式传递
package main

import "fmt"

func main() {
	var a, b int = 1, 2
	/*
	   调用 swap() 函数
	   &a 指向 a 指针，a 变量的地址
	   &b 指向 b 指针，b 变量的地址
	*/
	swap(&a, &b)

	fmt.Printf("a: %d b: %d \n", a, b)

	println(test("sum: %d", 1, 2, 3))

	s := []int{1, 2, 3}
	println(test("slice sum: %d", s...))
}

func swap(x, y *int) {
	var temp int = *x
	*x = *y
	*y = temp
}

// 不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）
func test(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}
	return fmt.Sprintf(s, x)
}
