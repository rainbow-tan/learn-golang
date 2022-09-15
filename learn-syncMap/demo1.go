package main

import (
	"fmt"
	"sync"
)

func traversal(k, v interface{}) bool {
	if k == "aaa" {
		fmt.Println("遍历到aaa, 结束遍历")
		return false
	} else {
		fmt.Printf("遍历到:%v, 继续遍历\n", k)
		return true
	}

}
func main() {
	var scene sync.Map
	// 将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))

	v, ok := scene.Load("no key")
	if !ok {
		fmt.Println("不存在键\"no key\"")
	} else {
		fmt.Printf("获取的值是:%v\n", v)
	}
	fmt.Println("===============================")
	// 根据键删除对应的键值对
	scene.Delete("london")
	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
	fmt.Println("===============================")
	scene.Store("aaa", 1)
	scene.Store("aaa", 2)
	scene.Store("bb", 3)

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate2:", k, v)
		return true
	})
	fmt.Println("===============================")
	// 遍历所有sync.Map中的键值对
	scene.Range(traversal)
}
