package main

import (
	"fmt"
	"net"
)

func main() {
	//启动本地服务
	listerner, err := net.Listen("tcp", "127.0.0.1:6001")
	if err != nil {
		fmt.Println("start server failed: ", err)
		return
	}
	for {
		//等待别人来连接
		conn, err := listerner.Accept()
		go func(conn net.Conn) {
			for {
				if err != nil {
					fmt.Println("Accept failed: ", err)
					return
				}
				//与客户端进行通信
				var buf [128]byte
				n, err := conn.Read(buf[:])
				if err != nil {
					fmt.Println("Read from conn failed: ", err)
					return
				}

				fmt.Printf("Read: [%s] \n", string(buf[:n]))
			}
		}(conn)
	}
}
