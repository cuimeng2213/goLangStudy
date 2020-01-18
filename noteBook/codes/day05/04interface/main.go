package main

import (
	"fmt"
)

//接口示例
//不管是什么牌子的车都能跑

type Car interface{
	run()
}

type falali struct{
	brand string
}
type baoshijie struct{
	brand string
}
func (f falali) run(){

	fmt.Printf("%s is run\n", f.brand)
}
func (f falali) stop(){
	fmt.Printf("%s is stop\n", f.brand)	
}
func (b baoshijie) run(){

	fmt.Printf("%s is run\n", b.brand)
}

//函数接收一个car的接口类型，不管是什么类型的结构体
//只要是实现run方法
func drive(car Car){
	car.run()
}

func main(){
	f1 := falali{
		brand:"法拉利",
	}
	b1 := baoshijie{
		brand:"保时捷",
	}
	drive(f1)
	drive(b1)
}