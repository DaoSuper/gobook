// Golang 表达式 ：根据调用者不同，方法分为两种表现形式:
// instance.method(args...) ---> <type>.func(instance, args...)
package main

import "fmt"

type User struct {
	id   int
	name string
}

func (u *User) TestPointer() {
	fmt.Printf("TestPointer: %p, %v\n", u, u)
}

func (u User) TestValue() {
	fmt.Printf("TestValue: %p, %v\n", &u, u)
}

func main() {
	u := User{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &u, u)

	mv := User.TestValue
	mv(u)

	mp := (*User).TestPointer
	mp(&u)

	mp2 := (*User).TestValue
	mp2(&u)

	var p *User = nil
	p.TestPointer2()

	(*User)(nil).TestPointer2() // method value
	(*User).TestPointer(nil)    // method expression   // method expression
}

// 函数
func (User) TestValue2() {}

func (*User) TestPointer2() {}
