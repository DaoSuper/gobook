// UDP协议（User Datagram Protocol） 用户数据报协议
// 是OSI（Open System Interconnection，开放式系统互联）参考模型中一种无连接的传输层协议，
// 不需要建立连接就能直接进行数据发送和接收，属于不可靠的、没有时序的通信
// 实时性比较好，通常用于视频直播相关领域

// UDP server端
package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer listen.Close()
	for {
		var data [1024]byte
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr) // 发送数据
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
	}

}
