package main

import (
	"fmt"
	"os"
)

// fmt包 主要分为向外输出内容和获取输入内容两大部分
func main() {
	fmt.Print("在终端打印该信息。")
	name := "枯藤"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")

	// 向标准输出写入内容
	fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./dao8/fmt_demo/fmt_demo.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	fname := "枯藤"
	// 向打开的文件句柄中写入内容
	fmt.Fprintf(fileObj, "往文件中写如信息：%s", fname)

	//Sprint系列函数会把传入的数据生成并返回一个字符串
	s1 := fmt.Sprint("枯藤")
	name3 := "枯藤"
	age := 18
	s2 := fmt.Sprintf("name:%s, age:%d", name3, age)
	s3 := fmt.Sprintln("枯藤")
	fmt.Println(s1, s2, s3)

	//通用占位符
	// %v	值的默认格式表示
	// %+v	类似%v，但输出结构体时会添加字段名
	// %#v	值的Go语法表示
	// %T	打印值的类型
	// %%	百分号
	fmt.Printf("%v\n", 100)
	fmt.Printf("%v\n", false)
	o := struct{ name string }{"枯藤"}
	fmt.Printf("%v\n", o)
	fmt.Printf("%#v\n", o)
	fmt.Printf("%T\n", o)
	fmt.Printf("100%%\n")

	// 布尔型
	//%t	true或false
	fmt.Printf("布尔型: %t\n", true)

	//整型
	// %b	表示为二进制
	// %c	该值对应的unicode码值
	// %d	表示为十进制
	// %o	表示为八进制
	// %x	表示为十六进制，使用a-f
	// %X	表示为十六进制，使用A-F
	// %U	表示为Unicode格式：U+1234，等价于”U+%04X”
	// %q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示

	n := 65
	fmt.Printf("二进制: %b\n", n)
	fmt.Printf("unicode码值: %c\n", n)
	fmt.Printf("十进制: %d\n", n)
	fmt.Printf("八进制: %o\n", n)
	fmt.Printf("十六进制，使用a-f: %x\n", n)
	fmt.Printf("为十六进制，使用A-F: %X\n", n)

	//浮点数与复数
	// %b	无小数部分、二进制指数的科学计数法，如-123456p-78
	// %e	科学计数法，如-1234.456e+78
	// %E	科学计数法，如-1234.456E+78
	// %f	有小数部分但无指数部分，如123.456
	// %F	等价于%f
	// %g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
	// %G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
	f := 12.34
	fmt.Printf("二进制指数的科学计数法: %b\n", f)
	fmt.Printf("科学计数法: %e\n", f)
	fmt.Printf("科学计数法: %E\n", f)
	fmt.Printf("有小数部分但无指数部分: %f\n", f)
	fmt.Printf("采用%%e或%%f格式: %g\n", f)
	fmt.Printf("采用%%E或%%F格式: %G\n", f)

	//字符串和[]byte
	// %s	直接输出字符串或者[]byte
	// %q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
	// %x	每个字节用两字符十六进制数表示（使用a-f
	// %X	每个字节用两字符十六进制数表示（使用A-F）
	s := "枯藤"
	fmt.Printf("直接输出: %s\n", s)
	fmt.Printf("字面值: %q\n", s)
	fmt.Printf("十六进制数（a-f): %x\n", s)
	fmt.Printf("十六进制数(A-F）: %X\n", s)

	//指针
	// %p	表示为十六进制，并加上前导的0x
	a := 18
	fmt.Printf("指针: %p\n", &a)
}
