// Golang匿名字段 ：可以像字段成员那样访问匿名字段方法，编译器负责查找。
package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
	title string
}

func (u *User) ToString() string {
	return fmt.Sprintf("User: %p, %v", u, u)
}

func (u *Manager) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", u, u)
}

func main() {
	m := Manager{User{1, "Tom"}, "Administrator"}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.User.ToString())

	fmt.Println(m.ToString())
}
