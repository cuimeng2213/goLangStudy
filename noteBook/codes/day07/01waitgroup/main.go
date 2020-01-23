package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func f1() {
	//指定随机种子
	now := time.Now().UnixNano()
	rand.Seed(now)
	r1 := rand.Int()
	r2 := rand.Intn(10)
	fmt.Println(r1, r2)
}

//
func sleep_demo(i int) {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
	fmt.Println("sleep_demo: ", i)
	wg.Done() // 计数器减一
}
func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1) //计数器加1
		go sleep_demo(i)
	}
	// 如何确定10个goroutine都结束呢？
	fmt.Println(runtime.NumCPU())
	wg.Wait() //等待wg的计数器减为0
}
