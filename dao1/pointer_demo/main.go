// &（取地址）和*（根据地址取值）
// 1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
// 2.指针变量的值是指针地址。
// 3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。
package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("a:%d ptr:%p\n", a, &a)
	fmt.Printf("b:%p type:%T\n", b, b)
	fmt.Println(&b)

	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	modify1(a)
	fmt.Println("value a:", a)
	modify2(&a)
	fmt.Println("value &a:", a)

	newTest()
	makeTest()

	sliceTest()
}

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

// 1.二者都是用来做内存分配的。
// 2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
// 3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
func newTest() {
	var a *int = new(int)
	*a = 10
	fmt.Println(*a)
}

func makeTest() {
	var b map[string]int = make(map[string]int, 10)
	b["测试"] = 100
	fmt.Println(b)
}

func sliceTest() {
	var a int
	fmt.Println(&a)
	var p *int = &a
	*p = 20
	fmt.Println(a)
}
