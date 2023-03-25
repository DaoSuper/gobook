// 接口（interface）定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
// 是一种类型，一种抽象的类型。
// interface是一组method的集合，是duck-type programming的一种体现。

/*
接口是一个或多个方法签名的集合。
任何类型的方法集中只要拥有该接口'对应的全部方法'签名。
就表示它 "实现" 了该接口，无须在该类型上显式声明实现了哪个接口。
这称为Structural Typing。
所谓对应方法，是指有相同名称、参数列表 (不包括参数名) 以及返回值。
当然，该类型还可以有其他方法。

接口只有方法声明，没有实现，没有数据字段。
接口可以匿名嵌入其他接口，或嵌入到结构中。
对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，既无法修改复制品的状态，也无法获取指针。
只有当接口存储的类型和对象都为nil时，接口才等于nil。
接口调用不会做receiver的自动转换。
接口同样支持匿名字段方法。
接口也可实现类似OOP中的多态。
空接口可以作为任何类型数据的容器。
一个类型可实现多个接口。
接口命名习惯以 er 结尾。
*/
package main

import "fmt"

func main() {

	var ser Sayer

	a := cat{}
	b := dog{}
	ser = a
	ser.say()
	ser = b
	ser.say()
}

type Sayer interface{ say() }

type dog struct{}
type cat struct{}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}
