package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read_file(s string) {
	fp, err := os.Open(s)
	if err != nil {
		fmt.Printf("read file %s failed\n", s)
		return
	}
	reader := bufio.NewReader(fp)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("%s", line)
			fmt.Printf("read over\n")
			return
		}
		fmt.Printf("%s", line)
	}
}

func main() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("%#v\n", err)
		return
	}
	var tmp = make([]byte, 128)
	//读取文件内容
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("read error：", err)
		return
	}
	fmt.Printf("read %d \n", n)
	fmt.Printf("data: %s \n", string(tmp[:n]))
	//关闭文件
	fileObj.Close()
	read_file("./main.go")
}
