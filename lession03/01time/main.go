package main

import (
	// 如果导入的包名与go文件中package声明的不一样， 则有如下行为： "cuimeng/lession03/01time/utils" as pkg1  也就是起别名
	// import pgk1 "cuimeng/lession03/01time/utils"
	"cuimeng/lession03/01time/utils"
	"fmt"
	"time"
)

// time包练习
func showDetail(t time.Time) {
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Unix())
}

func time2str(t time.Time) string {
	time.Sleep(30)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	now := time.Now()
	fmt.Printf("%T\n", now)
	showDetail(now)

	s := time2str(now)
	fmt.Println(s)
	x := pkg1.Add(1, 2)
	pkg1.ShowTime()
	fmt.Println(x)
}
