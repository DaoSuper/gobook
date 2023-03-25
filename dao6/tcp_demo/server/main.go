// TCP/IP(Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间协议，
// 是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（Transport layer）通信协议

// 服务端
// 1.监听端口
// 2.接收客户端请求建立链接
// 3.创建goroutine处理链接。
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

// 处理函数
func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(fmt.Sprintf("server已接受到: %v", recvStr))) // 发送数据
	}
}
