package main

import (
	"fmt"
	"os"
)

// 文件操作
func file_ops() {
	fp, err := os.Open("./main.go")
	// defer fp.Close()
	if err != nil {
		fmt.Println("Open file error: ", err)
		return
	}

	defer fp.Close()
}
func main() {
	var a interface{}
	a = 100
	switch a.(type) {
	case int32:
		fmt.Println("INT32")
	case int:
		fmt.Println("INT") //int 类型
	case string:
		fmt.Println("STRING")
	}
	file_ops()
}
