// 闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)。
package main

import "fmt"

func main() {
	c := a()
	c()
	c()
	c()

	a()

	f := test()
	f()
}

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

// 闭包复制的是原对象指针
func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}
