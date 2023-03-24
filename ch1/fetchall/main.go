//goroutine和channel
//同时去获取所有的URL，所以这个程序的总执行时间不会超过执行时间最长的那一个任务

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递
func main() {
	start := time.Now()
	ch := make(chan string)

	var urlList [3]string = [3]string{"https://juejin.cn/", "https://books.studygolang.com/gopl-zh/", "https://golang.org"}

	for _, url := range urlList {
		go fetch(url, ch) // start a goroutine
	}
	for range urlList {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
