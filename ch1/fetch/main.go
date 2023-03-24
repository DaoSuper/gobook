// 获取对应的url，并将其源文本打印出来
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	var url = "https://books.studygolang.com/gopl-zh/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("fetch: %v\n", err)
	}
	b, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("fetch: reading %s: %v\n", url, err)
	}
	fmt.Printf("%s", b)
}
