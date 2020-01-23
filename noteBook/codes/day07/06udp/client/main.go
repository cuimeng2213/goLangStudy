package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	sk, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 7001,
	})
	if err != nil {
		fmt.Println("DialUDP Error: ", err)
		return
	}
	defer sk.Close()
	var replay [1024]byte
	var reader = bufio.NewReader(os.Stdin)
	//发送数据
	for {
		fmt.Print("请输入内容： ")
		s, _ := reader.ReadString('\n')
		_, err := sk.Write([]byte(s))
		if err != nil {
			fmt.Println("Write Failed: ", err)
			return
		}

		n, _, err := sk.ReadFromUDP(replay[:])
		if err != nil {
			fmt.Println("ReadFromUDP Error: ", err)
			return
		}
		fmt.Println(string(replay[:n]))

	}
}
