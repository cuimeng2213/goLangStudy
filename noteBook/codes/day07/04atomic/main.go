package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var x int64

func add() {
	// 不加锁，或者不使用原子操作，得出的最终值不是预期的值。
	//x++
	//保证线程安全执行。
	atomic.AddInt64(&x, 1)
	wg.Done()
}
func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}
