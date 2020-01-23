package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type JobChan struct {
	num int64
}

type ResultChan struct {
	*JobChan
	sum  int64
	goId int
}

var wg sync.WaitGroup

var jobChan = make(chan *JobChan, 100)
var resultChan = make(chan *ResultChan, 100)

func producer(ch chan<- *JobChan) {
	rand.Seed(time.Now().UnixNano())
	defer close(ch)
	for {
		ch <- &JobChan{num: rand.Int63()}
	}
}

func worker(ch1 <-chan *JobChan, ch2 chan<- *ResultChan, n int) {
	// defer close(ch1)
	defer close(ch2)
	for {
		ret, ok := <-ch1
		if !ok {
			continue
		}
		//计算每个位累加
		tmp := ret.num
		var sum int64 = 0
		for tmp > 0 {
			sum += (tmp % 10)
			tmp /= 10
		}
		//构造ResultChan结构体
		r := &ResultChan{
			JobChan: ret,
			sum:     sum,
			goId:    n,
		}
		ch2 <- r
	}
}

func main() {
	//启动生产int64数值
	wg.Add(1)
	go producer(jobChan)
	//启动6个线程计算
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go worker(jobChan, resultChan, i)
	}
	//读取结果通道数据
	for {
		r, ok := <-resultChan
		if !ok {
			continue
		}
		fmt.Printf("goroutine=%d ### %d = %d \n", r.goId, r.JobChan.num, r.sum)
	}

	wg.Wait()

}
