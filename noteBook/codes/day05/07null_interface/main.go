package main
import (
	"fmt"
)
// 空接口
//interface 是关键字
// interface{} 这个是空接口类型

//空接口作为参数
func show( a interface{}){
	fmt.Printf("%T %v \n", a, a)
}

func main(){
	m1 := make(map[string]interface{}, 16)
	m1["name"]="cuimeng"
	m1["age"]=100
	m1["merried"]=true
	m1["hobby"]=[...]string{"唱歌","跳舞"}
	fmt.Printf("%#v\n", m1)
	show(true)
	show(nil)
}