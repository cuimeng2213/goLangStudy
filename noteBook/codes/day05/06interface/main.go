package main
//值接收者和指针接收者区别
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
// 方法使用值接收者实现接口所有方法
// func (c cat) move(){
// 	fmt.Println("猫在移动")
// }
// func (c cat) eat(s string){
// 	fmt.Printf("猫在吃 %s", s)
// }

// type chiken struct{
// 	feet int8
// }
// 方法使用指针接收者实现接口所有方法
func (c *cat) move(){
	fmt.Println("猫在移动")
}
func (c *cat) eat(s string){
	fmt.Printf("猫在吃 %s", s)
}

type chiken struct{
	feet int8
}
//
func (c chiken) move(){
	fmt.Println("鸡在移动")
}
func (c chiken) eat(s string){
	fmt.Printf("鸡在吃 %s", s)
}

func main(){
	var a1 animal
	c1 := cat{
		name:"tom",
		feet:4,
	}
	c2 := &cat{
		name:"kate",
		feet:4,
	}

	a1 = c1

	a1 = c2 // 可否正常赋值？？？
	fmt.Printf("%T\n",a1)
}