package main

import (
	"fmt"
	"net"
)

func main() {
	//根据域名返回IP
	address := "www.baidu.com"
	addr, err := net.ResolveIPAddr("ip", address)
	if err != nil {
		fmt.Printf("根据域名获取IP失败, e:%s\n", err)
	}
	fmt.Printf("域名 %s 的IP是 %s\n", address, addr)

	//检查IP地址格式是否有效 - 有效返回IP 无效返回nil
	s := "172.17.130.40"
	ip := net.ParseIP(s)
	if ip != nil {
		fmt.Printf("这是一个合法的IP:%s\n", s)
	} else {
		fmt.Printf("这不是一个合法的IP:%s\n", s)
	}

	//获取系统的网卡信息
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("获取系统的网卡信息失败, err:%s\n", err)
	}
	for _, inter := range interfaces {
		fmt.Printf("网卡信息是:%+v\n", inter)
	}
}
