package main

import "fmt"

// 接口例子

type Sayer interface {
	say()
}

type Mover interface {
	Move()
}

// 接口继承
type Animal interface {
	Sayer
	Mover
}

type Cat struct {
	name string
	age  int
}

func (c Cat) say() {
	fmt.Printf("%s is saying\n", c.name)
}

func (c Cat) Move() {
	fmt.Printf("%s is move\n", c.name)
}

// 空接口使用
func showType(i interface{}) { //空接口作为参数可以接收任意类型
	fmt.Printf("type= %T\n", i)
}

func main() {
	var animal Animal
	c := Cat{name: "xiaomi", age: 12}
	animal = c
	animal.say()
	animal.Move()

	a := 12
	showType(a)
	b := "this is a test"
	showType(b)
	// 空接口作为 map 的value值
	var stu = make(map[string]interface{}, 100)
	stu["nihao"] = 22.0
	stu["崔猛"] = 33
	fmt.Println(stu, len(stu))

	// 接口校验
	var x interface{}

	x = 12
	k, v := x.(string)
	fmt.Println(k, v)
}
