package main

import (
	"fmt"
	"runtime"
	"time"
)

// runtime.GOMAXPROCS()函数设置当前程序并发时占用的CPU逻辑核心数
// Go1.5版本之后，默认使用全部的CPU逻辑核心数

// 1.一个操作系统线程对应用户态多个goroutine。
// 2.go程序可以同时使用多个操作系统线程。
// 3.goroutine和OS线程是多对多的关系，即m:n。
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
