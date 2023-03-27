// channel是一种引用类型
// 遵循先入先出的规则，保证收发数据的顺序。

// var 变量 chan 元素类型
// 声明的通道后需要使用make函数初始化之后才能使用

package main

import "fmt"

func main() {

	var ch chan int
	fmt.Println("chan int", ch)

	// 通道有发送（send）、接收(receive）和关闭（close）三种操作。
	// 发送和接收都使用<-符号。
	ch2 := make(chan int) //创建无缓冲的通道
	go recv(ch2)          // 启用goroutine从通道接收值
	ch2 <- 10             // 发送

	ch3 := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch3 <- 10
	fmt.Println("发送成功", ch3)

	//如果你的管道不往里存值或者取值的时候一定记得关闭管道
	ch4 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch4 <- i
		}
		close(ch4)
	}()
	for {
		if data, ok := <-ch4; ok {
			fmt.Println("ch4 data: ", data)
		} else {
			break
		}
	}

	hasClose()
}

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

//判断一个通道是否被关闭
// for rang (常用)
// i, ok := <-ch1

func hasClose() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println("range ch2: ", i)
	}

}
