package main

import (
	"fmt"
)

const (
	withe = iota
	red
)

func main() {
	fmt.Println(withe, red)
	s := "hello1"
	byteArray := []byte(s)
	s1 := ""
	for i := len(byteArray) - 1; i >= 0; i-- {
		s1 = s1 + string(byteArray[i])
	}
	fmt.Println(s1)
	// method 2
	length := len(byteArray)

	for i := 0; i < length/2; i++ {
		byteArray[i], byteArray[length-1-i] = byteArray[length-1-i], byteArray[i]
	}
	s1 = string(byteArray)
	fmt.Println(s1)
	// 打印99乘法口诀
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d = %d \t", i, j, i*j)
		}
		fmt.Println("")
	}
	switch {
	case length >= 1:
		fmt.Printf("%s", "1")
	case length == 2:
		fmt.Println("sssss")
	}
	s2 := "I love 中国"
	for _, v := range s2 {
		fmt.Println(v)
		fmt.Printf("%c\n", v)
	}
	a := []int{0, 0, 0}
	a[1] = 10
	b := make([]int, 3) // make slice
	b[1] = 10
	// c := new([]int)
	//c[1] = 10 // invalid operation: c[1] (type *[]int does not support indexing)
	type data struct{ a int }
	var d = data{1234}
	var p *data
	p = &d
	fmt.Printf("%p, %T\n", p, p.a)

	// 判断 200 - 1000 之间的数字哪些是质数
	for i := 200; i <= 1000; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 { // 非质数
				flag = false
				break
			}
		}
		if flag == true {
			fmt.Printf("%d \t", i)
		}
	}
	// 位运算
	fmt.Printf("\n%b -- %b\n", 13, 3)
	fmt.Printf("\n%b --\n", 13|3)
}
