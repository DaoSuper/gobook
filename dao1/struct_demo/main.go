// type 类型名 struct {
// 	字段名 字段类型
// 	字段名 字段类型
// 	…
// }
// 1.类型名：标识自定义结构体的名称，在同一个包内不能重复。
// 2.字段名：表示结构体字段名。结构体中的字段名必须唯一。
// 3.字段类型：表示结构体字段的具体类型。

// 只有当结构体实例化时，才会真正地分配内存
// var 结构体实例 结构体类型

package main

import "fmt"

// 类型定义
type NewInt int

// 类型别名
type MyInt = int

type person struct {
	name string
	city string
	age  int8
}

func main() {
	var a NewInt
	var b MyInt

	fmt.Printf("type of a:%T\n", a)
	fmt.Printf("type of b:%T\n", b)

	var p1 person
	p1.name = "pprof.cn"
	p1.city = "北京"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%#v\n\n", user)

	//创建指针类型结构体
	var p2 = new(person)
	p2.name = "测试"
	p2.age = 18
	p2.city = "北京"
	fmt.Printf("p2=%#v\n\n", p2)

	p5 := person{
		name: "person5",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n\n", p5)

	p6 := &person{
		name: "person6",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n\n", p6)

	studentTest()

	p9 := newPerson("person9", "测试", 90)
	fmt.Printf("%#v\n", p9)

	p9.Dream()

}

//	func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//		函数体
//	}
//
// 1.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
// 2.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
// 3.方法名、参数列表、返回参数：具体格式与函数定义相同。

// NewPerson 构造函数
func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

// Dream Person做梦的方法
func (p person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n\n", p.name)
}

type student struct {
	name string
	age  int
}

func studentTest() {
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {

		m[stu.name] = &stu
	}
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
