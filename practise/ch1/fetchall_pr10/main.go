//练习 1.10： 找一个数据量比较大的网站，用本小节中的程序调研网站的缓存策略，
//对每个URL执行两遍请求，查看两次时间是否有较大的差别，并且每次获取到的响应内容
//是否一致，修改本节中的程序，将响应结果输出，以便于进行对比。

package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	ch1 := make(chan string)
	var urlList [3]string = [3]string{"https://juejin.cn/", "https://books.studygolang.com/gopl-zh/", "https://baidu.com"}

	for _, url := range urlList {
		go fetch(url, ch) // start a goroutine// 并发获取url
	}
	for range urlList {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	start = time.Now()
	for _, url := range urlList {
		go fetch(url, ch1) // start a goroutine
	}
	for range urlList {
		fmt.Println(<-ch1) // receive from channel ch 但是接受的时候是单线程，因为要方式没有接受完就结束那种情况
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

//第二次访问会比第一次快。因为有缓存
