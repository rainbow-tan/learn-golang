package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("读取文件出错, file:%s error:%s\n", filename, err)
		return
	}
	fmt.Printf("bytes:%s\n", bytes)

}
func main() {
	//读取不包含中文的字符是OK的
	readFile("normal.txt")
	fmt.Printf("***********************\n")

	//读取不存在的文件报错 The system cannot find the file specified.
	readFile("no_exist_file.txt")
	fmt.Printf("***********************\n")

	//读取包含中文的字符也是OK的
	readFile("contain_chinese.txt")
	fmt.Printf("***********************\n")
}
