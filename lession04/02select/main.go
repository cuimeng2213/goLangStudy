package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	id     int
	number int64
}

type Result struct {
	item *Item
	sum  int64
}

var numChan chan *Item
var sumChan chan *Result

func Producer(numChan chan *Item) {
	id := 0
	for {
		number := rand.Int63()
		id++
		//fmt.Println(">>> pro: ", number)
		tmp := &Item{
			id:     id,
			number: number,
		}
		numChan <- tmp
	}

}

func Consumer(numChan chan *Item, sumChan chan *Result) {

	for {
		var sum int64 = 0
		ret := <-numChan
		tmpNum := ret.number
		//fmt.Printf(">> sum a number: %d\n", ret.number)
		for {

			if tmpNum == 0 {

				break
			}
			//fmt.Printf("sum a number: %d\n", sum)
			sum = tmpNum%10 + sum
			tmpNum = tmpNum / 10
		}

		sumChan <- &Result{
			item: ret,
			sum:  sum,
		}
	}

}

func callWorker(n int, numChan chan *Item, sumChan chan *Result) {
	for i := 0; i < n; i++ {
		go Consumer(numChan, sumChan)
	}
}

func printResult(sumChan chan *Result) {
	for ret := range sumChan {
		fmt.Printf(">>>: ret= sum=%d id=%d number=%d\n", ret.sum, ret.item.id, ret.item.number)
		time.Sleep(time.Second)
	}
}

func main() {
	numChan = make(chan *Item, 100)
	sumChan = make(chan *Result, 100)

	exitChan := make(chan struct{}, 1)

	go Producer(numChan)
	callWorker(20, numChan, sumChan)

	go printResult(sumChan)

	//开一个读取标准输入得匿名函数协程
	go func(ch chan struct{}) {
		for {
			b := make([]byte, 1)
			fmt.Printf("等待任意字符退出\n")
			fmt.Scan(&b)
			ch <- struct{}{}
		}
	}(exitChan)

	for {
		select {
		case <-exitChan:
			fmt.Println(">>>>>>>>>>>>>>>>>>> tuichule")
			// 此处break只能退出select语句块， 若要退出for循环则使用return
			return
		}
	}
}
