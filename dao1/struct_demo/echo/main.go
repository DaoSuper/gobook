package main

import "fmt"

type student struct {
	id   int
	name string
	age  int
}

func demo(ce []student) {
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
	// ce = append(ce, student{3, "xiaowang", 56})
	// return ce
}
func main() {
	var ce []student //定义一个切片类型的结构体
	ce = []student{
		{1, "xiaoming", 22},
		{2, "xiaozhang", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)

	// 删除map类型的结构体
	de := make(map[int]student)
	de[1] = student{1, "xiaolizi", 22}
	de[2] = student{2, "wang", 23}
	fmt.Println(de)
	delete(de, 2)
	fmt.Println(ce)
}
