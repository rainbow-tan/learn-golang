package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	//连接给定的network地址
	address := "127.0.0.1:8080"
	conn, err := net.Dial("udp", address)
	if err != nil {
		fmt.Println(fmt.Sprintf("连接失败, err:%s", err))
		return
	}
	fmt.Println(fmt.Sprintf("连接成功, address:%s", address))

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("关闭连接失败, err:%s", err))
		}
		fmt.Println(fmt.Sprintf("关闭连接成功"))
	}(conn)

	//写数据到UDP
	msg := fmt.Sprintf("我是来自于client的信息,现在的时间是:%s", time.Now().Format("2006-01-02 15:04:05"))
	msgByte := []byte(msg)
	n, err := conn.Write(msgByte)
	if err != nil {
		fmt.Println(fmt.Sprintf("client发送失败, err:%s", err))
		return
	}
	fmt.Println(fmt.Sprintf("client发送成功%d个字节", n))

}
