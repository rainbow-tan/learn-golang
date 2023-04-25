package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type MyStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

func readFileToStruct(filename string) {
	//读取文件转为结构体
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("读取文件出错, file:%s error:%s\n", filename, err)
		return
	}
	fmt.Printf("bytes:%s\n", bytes)

	var my MyStruct
	err = json.Unmarshal(bytes, &my)
	if err != nil {
		fmt.Printf("绑定到结构体出错, file:%s error:%s\n", filename, err)
		return
	}
	fmt.Printf("MyStruct:%+v\n", my)

}
func main() {
	//正常转换
	readFileToStruct("normal.json")
	fmt.Printf("***********************\n")

	//不存在的文件 读取时抛出异常 The system cannot find the file specified.
	readFileToStruct("no_file.json")
	fmt.Printf("***********************\n")

	//json key不区分关键字大小写
	readFileToStruct("case_insensitive.json")
	fmt.Printf("***********************\n")

	//json 少一个key 会绑定存在的key
	readFileToStruct("few_key.json")
	fmt.Printf("***********************\n")

	//json 多一个key 会绑定存在的key
	readFileToStruct("more_key.json")
	fmt.Printf("***********************\n")

	//json 格式错误 报错 invalid character '}' looking for beginning of object key string
	readFileToStruct("error_format.json")
	fmt.Printf("***********************\n")
}
