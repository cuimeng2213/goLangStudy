package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 7001,
	})
	if err != nil {
		fmt.Printf("Upd start failed %v \n", err)
		return
	}
	//不需要连接。直接收发数据
	var data [1024]byte
	for {
		n, clientAddr, err := conn.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("Read from failed: ", err)
			return
		}
		fmt.Println(data[:n])
		replay := strings.ToUpper(string(data[:n]))
		conn.WriteToUDP([]byte(replay), clientAddr)
	}
	conn.Close()
}
