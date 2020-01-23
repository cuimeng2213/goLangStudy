package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// tcp client
var reader = bufio.NewReader(os.Stdin)

func main() {
	//1.与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:6001")
	if err != nil {
		fmt.Println("Dial 127.0.0.1 failed")
		return
	}
	for {
		fmt.Print("请输入: ")
		s, err := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		if err != nil {
			fmt.Println("read stdin err: ", err)
			return
		}
		conn.Write([]byte(s))
	}
	conn.Close()
}
