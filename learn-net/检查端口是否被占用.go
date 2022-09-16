package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

// ScanPort 检查端口是否被占用
func ScanPort(network string, hostname string, port int) bool {
	p := strconv.Itoa(port)
	addr := net.JoinHostPort(hostname, p)

	//连接该端口 能连接到就是已经被占用了 连接不到就是没被占用
	conn, err := net.DialTimeout(network, addr, 3*time.Second)
	if err != nil {
		fmt.Println(fmt.Sprintf("该端口未被占用, %s:%d, err:%s", hostname, port, err))
		return false
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("关闭连接失败:%s", err))
		}
		fmt.Println("关闭连接成功")
	}(conn)
	fmt.Println(fmt.Sprintf("该端口已被占用, %s:%d", hostname, port))
	return true
}
func main() {
	used := ScanPort("tcp", "127.0.0.1", 8000)
	fmt.Println(fmt.Sprintf("是否在使用:%v", used))
}
