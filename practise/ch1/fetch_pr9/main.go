// 练习 1.9： 修改fetch打印出HTTP协议的状态码，可以从resp.Status变量得到该状态码。
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := "https://books.studygolang.com/gopl-zh/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch: %v\n", err)
	}
	w, err := io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("fetch: reading %s: %v\n", url, err)
	}
	fmt.Println("打印一下copy的字节数", w)
	fmt.Println("打印一下 HTTP协议状态码：", resp.Status)
}
