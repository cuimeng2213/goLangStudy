package main

import (
	"fmt"
)
// 定义一个animal 接口类型
type animal interface{
	move()
	eat(s string)
}

type cat struct{
	name string
	feet int8
}

func (c cat) move(){
	fmt.Println("猫在移动")
}
func (c cat) eat(s string){
	fmt.Printf("猫在吃 %s", s)
}

type chiken struct{
	feet int8
}

func (c chiken) move(){
	fmt.Println("鸡在移动")
}
func (c chiken) eat(s string){
	fmt.Printf("鸡在吃 %s", s)
}

func main(){
	var a1 animal
	a1 = cat{
		name:"kaite",
		feet:4,
	}
	a1.move()
	a1.eat("臭鱼")
	fmt.Printf("%T\n", a1) // main.cat
}