// • 只能为当前包内命名类型定义方法。
// • 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
// • 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
// • 不支持方法重载，receiver 只是参数签名的组成部分。
// • 可用实例 value 或 pointer 调用全部方法，编译器自动转换。
//方法是一种特殊的函数

// func (recevier type) methodName(参数列表)(返回值列表){}
package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func (u User) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

// 当接受者是指针时，即使用值类型调用那么函数内部也是对指针的操作
func (u *User) Notify2() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

func main() {
	// 值类型调用方法
	u1 := User{"golang", "golang@golang.com"}
	u1.Notify()
	u1.Notify2()

	// 指针类型调用方法
	u2 := User{"go", "go@go.com"}
	u3 := &u2
	u3.Notify()
	u3.Notify2()
}
