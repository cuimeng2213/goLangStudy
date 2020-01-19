package main

import (
	"fmt"
	"os"
	"time"
)

func timer() {
	tick := time.Tick(time.Second)
	for t := range tick {
		fmt.Println(t)
	}
}

func main() {
	t := time.Now()
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Date())
	fmt.Println(t.Hour())
	fmt.Fprintf(os.Stdout, "==%s==\n", "hello funck")
	// 格式化时间
	// 将go语言中的时间对象转换为字符串类型
	// 2019-08-03
	fmt.Println(t.Format("2006-01-02"))
	fmt.Println(t.Format("2006/01/02 15:04:05"))
	fmt.Println(t.Format("2006/01/02 15:04:05.000"))
	fmt.Println(t.Format("2006/01/02 15:04:05.9999"))

	//按照对应时间格式解析字符串类型的时间
	timeObj, err := time.Parse("2006-01-02 15:04:05", "2020-01-19 12:30:22")
	if err != nil {
		fmt.Println("Parse failed")
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Unix())
	fmt.Println(timeObj.Local().Sub(t.Local()))
	loc, _ := time.LoadLocation("Asia/Shanghai")
	t1, err := time.ParseInLocation("2006-01-02 15:04:05", "2020-01-19 21:13:00", loc)
	if err != nil {
		fmt.Println("####: ", err)
		return
	}
	fmt.Println("####: ", t1)
	timer()
}
