package main

import (
	"fmt"
	"net"
)

func main() {

	//把地址转换为UDP的地址
	address := "127.0.0.1:8080"
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(fmt.Sprintf("转换UDP地址失败, err:%s", err))
		return
	}
	fmt.Println(fmt.Sprintf("转换UDP地址成功, address:%s", address))

	//监听UDP
	udp, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println(fmt.Sprintf("监听UDP失败, err:%s", err))
		return
	}
	fmt.Println(fmt.Sprintf("监听UDP成功, address:%s", address))

	//从UDP中读取数据
	b := make([]byte, 1024)
	n, addr, err := udp.ReadFromUDP(b)
	if err != nil {
		fmt.Println(fmt.Sprintf("接收客户端的数据失败, err:%s", err))
		return
	}
	fmt.Println(fmt.Sprintf("接收客户端的数据成功, 接收到的字节数:%d, data:%s", n, b))
	fmt.Println(fmt.Sprintf("客户端的信息:%+v", addr.String()))

}
