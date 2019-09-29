package main

import "fmt"

var (
	name   string
	age    int
	gender string
)

type student struct {
	name   string
	age    int8
	gender string
	hobby  []string
}

type nodeDom struct {
	name string
	next *nodeDom
}

// 定义一个初始化函数
func newStu(name string, age int8, gender string, hobby []string) *student {
	return &student{
		name:   name,
		age:    age,
		gender: gender,
		hobby:  hobby,
	}
}
func newNode(name string, node *nodeDom) *nodeDom {
	return &nodeDom{
		name: name,
		next: node,
	}
}
func testSlice() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := s[2:5]
	s2 := s1[2:6:7]
	fmt.Printf("s1 %d %d %v \n", len(s1), cap(s1), s1)
	fmt.Printf("s1 %d %d %v \n", len(s2), cap(s2), s2)
}
func main() {
	s := "中国"
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
	for _, v := range s {
		fmt.Printf("%c\n", v)
	}
	//fmt.Scan(&name, &age, &gender)
	stu01 := newStu("张三", 18, "男", []string{"篮球", "羽毛球"})
	fmt.Printf("type= %T %s\n", stu01, stu01.name)

	fmt.Scanln(&name, &age, &gender)
	fmt.Println(name, age, gender)

	// test node
	n1 := newNode("n1", nil)
	fmt.Printf("n1: %T %s %p\n", n1, n1.name, n1.next)

	n2 := newNode("n2", n1)
	fmt.Printf("n2: %T %s %p\n", n2, n2.name, n2.next)
	testSlice()
}
