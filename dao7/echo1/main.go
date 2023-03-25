package main

import (
	"fmt"
	"sync"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}
func main() {
	go hello()
	fmt.Println("main goroutine done!")

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go multHello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
	time.Sleep(time.Second)

}

var wg sync.WaitGroup

// 启动多个goroutine
func multHello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
