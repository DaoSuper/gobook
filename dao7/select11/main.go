package main

import (
	"fmt"
	"time"
)

//select实现从多个通道接收值
// select {
// case <-chan1:
//    // 如果chan1成功读到数据，则进行该case处理语句
// case chan2 <- 1:
//    // 如果成功向chan2写入数据，则进行该case处理语句
// default:
//    // 如果上面都没有成功，则进入default处理流程
// }

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)

	go test1(output1)
	go test2(output2)
	// 用select监控
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}

	test3()
	test4()
	fmt.Println("main结束")
}

// 如果多个channel同时ready，则随机选择一个执行
func test3() {
	// 创建2个管道
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	go func() {
		//time.Sleep(2 * time.Second)
		int_chan <- 1
	}()
	go func() {
		string_chan <- "hello"
	}()
	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	}
}

// 可以用于判断管道是否存满
func test4() {
	// 创建管道
	output1 := make(chan string, 10)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
