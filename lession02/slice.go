package main

import "fmt"

func sum_target(arr []int, target int) {
	for index, v := range arr {
		other := target - v
		for i := index + 1; i < len(arr); i++ {
			if arr[i] == other {
				fmt.Printf("(%d %d)\n", index, i)
			}
		}
	}
}

func main() {
	// 初始化数组
	a := [...]int{1, 3, 5, 7, 9, 11, 13}
	g := a[:2]
	// 通过数组得到切片b
	b := a[:]
	b[0] = 100
	fmt.Println(a[0], b[0])

	c := b[2:5]
	fmt.Println(c, len(c), cap(c)) // [5 7 9 ] len=3 cap =5
	fmt.Printf("c: %p\n", c)
	// 切片越界了
	// d := c[:6]
	// fmt.Printf("%v\n", d)

	d := c[:5]
	fmt.Printf("%v %p \n", d, d)

	e := c[2:5]
	fmt.Printf("%v %p cap=%d\n", e, e, cap(e))
	// 通过append扩容 e地址发生变化， 切片扩容每次是前一次的2倍， 3*2 = 6
	// e = append(e, 100,200,300)
	// 如果数据大于一次的扩容，则以当前切片容器长度 + 扩容数据大小 3 + 9 = 12
	e = append(e, 100, 200, 300, 4, 5, 6, 7, 89, 90)
	fmt.Printf("%v %p cap=%d \n", e, e, cap(e))

	f := []int{1, 3, 5, 7, 8, 11}
	sum_target(f, 8)
}
