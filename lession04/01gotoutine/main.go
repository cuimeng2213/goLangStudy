package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	fmt.Println("hello")
	// 等待输入结束，确保goroutine执行结束
	var a int
	fmt.Scan(&a)
}
