package main

import (
	"fmt"
)

// 类型断言

func assign(a interface{}){
	s, ok := a.(string) // 尝试将a转换为string类型
	if !ok{
		fmt.Printf("assign failed\n")
	} else{
		fmt.Printf("%s\n",s)	
	}
	
}

//使用switch进行类型判断
func justfy(a interface{}){
	switch a.(type){
	case string:
		fmt.Printf("Is a string %s\n",a.(string))
	case int:
		fmt.Printf("Is a int %d\n",a.(int))
	}
}

func main(){
	assign(100)
	assign("hello")
	justfy("hello cuimeng")
	justfy(123)
}