// WebSocket是一种在单个TCP连接上进行全双工通信的协议
// WebSocket使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据
// 在WebSocket API中，浏览器和服务器只需要完成一次握手，两者之间就直接可以创建持久性的连接，并进行双向数据传输
// 需要安装第三方包：
// cmd中：go get -u -v github.com/gorilla/websocket
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}

}
