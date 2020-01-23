package main

import (
	"fmt"
	"sync"
)

func f1(ch1 chan<- int) {
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
	wg.Done()
}

//单向通道
//ch1 <-chan int ：只读channel
//ch2 chan<- int ：只写通道
func f2(ch1 <-chan int, ch2 chan<- int) {
	// for x := range ch1 {
	// 	ch2 <- (x * x)
	// }
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	close(ch2)
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	a := make(chan int, 100)
	b := make(chan int, 100)
	//1.启动一个goroutine，生成100个数发送到cha1中
	go f1(a)
	//2.启动一个goroutine，从ch1中取值，计算平方放到ch2中
	go f2(a, b)
	//3.在main函数中 重ch2中取值打印。
	for x := range b {
		fmt.Printf("%T %d \n", x, x)
	}
	wg.Wait()
}
