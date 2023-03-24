// 延迟调用（defer）
// 1. 关键字 defer 用于注册延迟调用。
// 2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
// 3. 多个defer语句，按先进后出的方式执行。
// 4. defer语句中的变量，在defer声明时就决定了。

// 1. 关闭文件句柄
// 2. 锁资源释放
// 3. 数据库连接释放
package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}

func main() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}

	ts := []Test{{"a"}, {"b"}, {"c"}}

	for _, t := range ts {
		t2 := t
		defer t2.Close()
	}

	test()

	deferToReturn()
}

// 延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取
func test() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y)
	}(x)

	x += 10
	y += 100
	println("x: ", x, "y: ", y)
}

// 在有具名返回值的函数中（这里具名返回值为 i），执行 return 2 的时候实际上已经将 i 的值重新赋值为 2
func deferToReturn() (i int) {
	i = 0
	defer func() {
		fmt.Println("deferToReturn", i)
	}()

	return 2
}
